package src

import (
	"context"
	"fmt"
	"math"
	"os"
	"sync"
	"time"

	"github.com/AlyxPink/minecraft-screenshot-bot/src/ai"
	"github.com/AlyxPink/minecraft-screenshot-bot/src/minecraft"
	"github.com/AlyxPink/minecraft-screenshot-bot/src/screenshot"
	"github.com/AlyxPink/minecraft-screenshot-bot/src/uploader"
	"github.com/charmbracelet/log"
)

const (
	// Check Mastodon limits: https://github.com/mastodon/mastodon/blob/e8605a69d22e369e34914548338c15c053db9667/app/models/scheduled_status.rb#L16-L17
	SHOTS = 24
)

func Start() {
	var wg sync.WaitGroup // Initialize a WaitGroup

	minecraft.Launch()
	minecraft.CreateNewWorld()
	minecraft.SetupScreenshot()
	for i := 0; i < SHOTS; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		log.Info(fmt.Sprintf("** Taking a new screenshot (%d/%d) **", i, SHOTS))
		minecraft.SetRandomTime()
		minecraft.SetRandomWeather()
		minecraft.TeleportPlayer()
		minecraft.TakeRandomScreenshot()
		latestScreenshot := screenshot.GetLatestScreenshot()
		log.Info(fmt.Sprintf("Screenshot path: %s", latestScreenshot.Path))
		wg.Add(1) // Indicate that there's one more goroutine to wait for
		go func(i int) {
			defer cancel()
			defer wg.Done() // Signal that this goroutine is done
			if os.Getenv("DRY_RUN") != "" {
				log.Warn("DRY_RUN is set, skipping uploads")
				return
			}
			uploadScreenshot(ctx, i, latestScreenshot)
		}(i)
	}
	minecraft.QuitGame()
	cleanup()
	wg.Wait() // Wait for all goroutines to finish
}

func uploadScreenshot(ctx context.Context, i int, s screenshot.Screenshot) {
	log.SetPrefix(fmt.Sprintf("Screenshot ID: %s", s.ID.String()))
	// Set up the upload struct
	upload := uploader.Upload{Screenshot: s}

	// Upload to R2 first, so OpenAI can get the URL to describe the image
	r2 := &uploader.R2{}
	err, url := r2.Upload(ctx, upload)
	if err == nil {
		// Use the R2 URL to get alt text from OpenAI
		altText := ai.DescribeImage(ctx, url)
		upload.Screenshot.AltText = altText
	} else {
		log.Error("Error uploading to R2: %v", err)
		log.Warn("Skipping getting alt text from OpenAI")
	}

	// Load the uploaders
	uploaders := []uploader.Uploader{
		&uploader.Mastodon{Iteration: i},
	}
	// Dispatch the uploads
	for _, u := range uploaders {
		uploadWithRetry(ctx, u, upload)
	}
}

func uploadWithRetry(ctx context.Context, u uploader.Uploader, upload uploader.Upload) {
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		err, url := u.Upload(ctx, upload)
		if err == nil {
			log.Info(fmt.Sprintf("Uploaded to %T: %s", u, url))
			break
		}

		retryDelay = time.Duration(float64(retryDelay) * math.Pow(2, float64(i)))
		log.Error("Error uploading", "retryDelay", retryDelay, "upload", upload, "err", err)

		// Sleep for the current delay, then double it for the next iteration
		time.Sleep(retryDelay)
	}
}

func cleanup() {
	// TODO: delete latest created world
	// TODO: delete screenshots files (?)
}

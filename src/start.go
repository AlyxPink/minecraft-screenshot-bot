package src

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/VictorBersy/minecraft-screenshot-bot/src/ai"
	"github.com/VictorBersy/minecraft-screenshot-bot/src/screenshot"
	"github.com/VictorBersy/minecraft-screenshot-bot/src/uploader"
	"github.com/charmbracelet/log"
)

const (
	SHOTS = 24
)

func Start() {
	// launchGame()
	// createNewWorld()
	// setupScreenshot()
	for i := 0; i < SHOTS; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		// 	log.Info(fmt.Sprintf("** Taking a new screenshot (%d/%d) **", i, SHOTS))
		// 	setRandomTime()
		// 	setRandomWeather()
		// 	teleportPlayer()
		// 	takeRandomScreenshot()
		latestScreenshot := screenshot.GetLatestScreenshot()
		log.SetPrefix(fmt.Sprintf("Screenshot ID: %s", latestScreenshot.ID))
		go uploadScreenshot(ctx, i, latestScreenshot)
	}
	// quitGame()
	// cleanup()
}

func uploadScreenshot(ctx context.Context, i int, s screenshot.Screenshot) {
	// Set up the upload struct
	upload := uploader.Upload{Screenshot: s}

	// Upload to R2 first, so OpenAI can get the URL to describe the image
	r2 := &uploader.R2{}
	err, url := r2.Upload(ctx, upload)
	if err != nil {
		log.Error("Error uploading to R2: %v", err)
		log.Warn("Skipping getting alt text from OpenAI")
		log.Warn("Skipping Mastodon upload")
		return
	}

	// Use the R2 URL to get alt text from OpenAI
	altText := ai.DescribeImage(ctx, url)
	upload.Screenshot.AltText = altText

	// Prepare the uploaders
	uploaders := []uploader.Uploader{
		&uploader.Mastodon{Iteration: i},
	}
	// Dispatch the uploads
	for _, u := range uploaders {
		go uploadWithRetry(u, upload)
	}
}

func uploadWithRetry(u uploader.Uploader, upload uploader.Upload) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		err, url := u.Upload(ctx, upload)
		if err == nil {
			log.Info("Uploaded to %T: %s", u, url)
			continue
		}

		log.Error("Error uploading to %T: %v", u, err)
		// Sleep for the current delay, then double it for the next iteration
		time.Sleep(retryDelay)
		retryDelay = time.Duration(float64(retryDelay) * math.Pow(2, float64(i)))
	}
}

func cleanup() {
	// TODO: delete latest created world
	// TODO: delete screenshots files (?)
}

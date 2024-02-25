package src

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/VictorBersy/minecraft-screenshot-bot/src/describer"
	"github.com/VictorBersy/minecraft-screenshot-bot/src/minecraft"
	"github.com/VictorBersy/minecraft-screenshot-bot/src/screenshot"
	"github.com/VictorBersy/minecraft-screenshot-bot/src/uploader"
	"github.com/charmbracelet/log"
)

const (
	SHOTS = 24
)

func Launch() {
	startBatch(SHOTS)
}

func startBatch(batchSize int) {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	log.Info(fmt.Sprintf("Starting batch of %d screenshots", batchSize))

	minecraft := minecraft.New()
	minecraft.Setup()

	for i := 0; i < batchSize; i++ {
		wg.Add(1)
		// Setups the game and takes a screenshot
		s := minecraft.CaptureScreenshot()

		// Upload the screenshot
		go func(i int) {
			startUploadTask(i, s)
			defer wg.Done() // Signal that this goroutine is done
		}(i)
	}

	minecraft.Quit()
	wg.Wait() // Wait for all goroutines to finish
}

func startUploadTask(i int, s screenshot.Screenshot) {
	// Set up the context and logger
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: fmt.Sprintf("[%d/%d] Screenshot ID: %s", i+1, SHOTS, s.ID.String()),
	})
	ctx = log.WithContext(ctx, logger)

	// Start upload process
	// TODO: Allow the user to choose the uploaders
	upload := uploader.Upload{Screenshot: s}
	publicURL, err := upload.UploadWithRetry(ctx, &uploader.R2{})
	if err != nil {
		log.FromContext(ctx).Error("Error uploading to R2, skip the rest of the upload", "err", err)
		return
	}

	// Get alt text using public upload URL
	describer := describer.Get()
	s.Description, err = describer.GenerateFromURL(ctx, publicURL)
	if err != nil {
		log.FromContext(ctx).Error("Error whle getting image description, skip the rest of the upload", "err", err)
		return
	}

	// Upload to Mastodon
	_, err = upload.UploadWithRetry(ctx, &uploader.Mastodon{Iteration: i})
	if err != nil {
		log.FromContext(ctx).Error("Error uploading to Mastodon", "err", err)
		return
	}
}

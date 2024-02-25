package uploader

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/VictorBersy/minecraft-screenshot-bot/src/screenshot"
	"github.com/charmbracelet/log"
)

type Upload struct {
	Screenshot *screenshot.Screenshot
}

type Uploader interface {
	upload(ctx context.Context, u Upload) (err error, url string)
}

func (upload Upload) UploadWithRetry(ctx context.Context, u Uploader) (url string, err error) {
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		err, url := u.upload(ctx, upload)

		if err != nil {
			// Exponential backoff
			retryDelay = time.Duration(float64(retryDelay) * math.Pow(2, float64(i)))

			log.FromContext(ctx).Error("Error during upload attempt", "uploader", u, "err", err)
			time.Sleep(retryDelay)
			continue
		}

		log.FromContext(ctx).Info(fmt.Sprintf("Uploaded to %T: %s", u, url))
		return url, err
	}

	return "", fmt.Errorf("Failed to upload after %d attempts", maxRetries)
}

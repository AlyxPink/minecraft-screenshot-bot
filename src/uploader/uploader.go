package uploader

import (
	"context"

	"github.com/AlyxPink/minecraft-screenshot-bot/src/screenshot"
)

type Upload struct {
	Screenshot screenshot.Screenshot
}

type Uploader interface {
	Upload(ctx context.Context, u Upload) (err error, url string)
}

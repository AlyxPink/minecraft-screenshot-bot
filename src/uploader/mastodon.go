package uploader

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/mattn/go-mastodon"
)

type Mastodon struct {
	Iteration int
}

func (u Mastodon) Upload(ctx context.Context, upload Upload) (error, string) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	// Upload media to Mastodon
	var attachment *mastodon.Attachment
	attachment, err := c.UploadMediaFromMedia(ctx, &mastodon.Media{
		File:        bytes.NewReader(upload.Screenshot.File),
		Description: upload.Screenshot.AltText.Long,
	})
	if err != nil {
		log.FromContext(ctx).Error("Error while uploading screenshot", "screenshot ID", upload.Screenshot.ID, "error", err)
		return err, ""
	}

	// Schedule post
	scheduledAt := time.Now().Add(time.Hour * 4 * time.Duration(u.Iteration))

	post := &mastodon.Toot{
		MediaIDs:    []mastodon.ID{attachment.ID},
		Sensitive:   false,
		Visibility:  mastodon.VisibilityUnlisted,
		Language:    "EN",
		ScheduledAt: &scheduledAt,
	}

	status, err := c.PostStatus(ctx, post)
	if err != nil {
		log.FromContext(ctx).Fatal(err)
	}

	log.FromContext(ctx).Info("Post scheduled", "scheduledAt", scheduledAt.String(), "statusID", status.ID)

	return nil, attachment.URL
}

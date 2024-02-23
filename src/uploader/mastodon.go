package uploader

import (
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
	log.SetPrefix("Mastodon uploader")

	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	// Upload media to Mastodon
	var attachment *mastodon.Attachment
	attachment, err := c.UploadMediaFromMedia(ctx, &mastodon.Media{
		File:        upload.Screenshot.File,
		Description: upload.Screenshot.AltText.Long,
	})
	if err != nil {
		log.Fatal("Error while uploading screenshot", "screenshot ID", upload.Screenshot.ID, "error", err)
	}
	log.Info("Screenshot uploaded", "screenshot ID", upload.Screenshot.ID, "attachment URL", attachment.URL)

	// Schedule post
	scheduledAt := time.Now().Add(time.Hour * 4 * time.Duration(u.Iteration)) // TODO: Set to X hours after latest post

	post := &mastodon.Toot{
		MediaIDs:  []mastodon.ID{attachment.ID},
		Sensitive: false,
		//Visibility:  "unlisted", // TODO: edit
		Visibility:  mastodon.VisibilityDirectMessage, // TODO: edit
		Language:    "EN",
		ScheduledAt: &scheduledAt,
	}

	status, err := c.PostStatus(ctx, post)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Post scheduled", "scheduledAt", scheduledAt.String(), "status ID", status.ID)

	return nil, attachment.URL
}

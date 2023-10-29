package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/mattn/go-mastodon"
)

func postScreenshotToSocialMedia(screenFile string, iteration int) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	media, err := c.UploadMedia(context.Background(), screenFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(fmt.Sprintf("Screenshot uploaded %s (%s)", screenFile, media.URL))

	scheduledAt := time.Now().UTC().Add(time.Hour * time.Duration(iteration)) // TODO: Set to X hours after latest post
	status, err := c.PostStatus(context.Background(), &mastodon.Toot{
		MediaIDs:    []mastodon.ID{media.ID},
		Sensitive:   false,
		Visibility:  "unlisted",
		Language:    "EN",
		ScheduledAt: &scheduledAt,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(fmt.Sprintf("Post scheduled at %s (%s)", scheduledAt.String(), status.URL))
}

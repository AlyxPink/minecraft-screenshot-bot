package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/mattn/go-mastodon"
)

func postScreenshotToSocialMedia(screenshot *os.File, iteration int) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	// Upload media to Mastodon
	var attachment *mastodon.Attachment
	attachment, err := c.UploadMediaFromMedia(context.Background(), &mastodon.Media{
		File:        screenshot,
		Description: "Randomly generated Minecraft screenshot by CraftViews bot.", // Default alt-text
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(fmt.Sprintf("Screenshot uploaded %s (%s)", screenshot.Name(), attachment.URL))

	// Schedule post
	scheduledAt := time.Now().Add(time.Hour * 4 * time.Duration(iteration)) // TODO: Set to X hours after latest post
	post := &mastodon.Toot{
		MediaIDs:    []mastodon.ID{attachment.ID},
		Sensitive:   false,
		Visibility:  "unlisted",
		Language:    "EN",
		ScheduledAt: &scheduledAt,
	}
	status, err := c.PostStatus(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(fmt.Sprintf("Post scheduled at %s (ID: %s)", scheduledAt.String(), status.ID))
}

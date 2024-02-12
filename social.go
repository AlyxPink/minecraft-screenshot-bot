package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/mattn/go-mastodon"
)

func postScreenshotToSocialMedia(screenshot_path string, screenshot_url url.URL, iteration int) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	// Open screenshot file
	screenshot, err := os.Open(screenshot_path)
	if err != nil {
		log.Fatal(err)
	}

	// Get alt-text
	altText := describeImage(screenshot_url)

	// Upload media to Mastodon
	var attachment *mastodon.Attachment
	attachment, err = c.UploadMediaFromMedia(context.TODO(), &mastodon.Media{
		File:        screenshot,
		Description: altText,
	})
	log.Info("File", "screenshot", screenshot)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while uploading screenshot to Mastodon %s (%s)", screenshot.Name(), err))
	}
	log.Info(fmt.Sprintf("Screenshot uploaded %s (%s)", screenshot.Name(), attachment.URL))

	// Schedule post
	//scheduledAt := time.Now().Add(time.Hour * 4 * time.Duration(iteration)) // TODO: Set to X hours after latest post
	scheduledAt := time.Now()
	post := &mastodon.Toot{
		MediaIDs:  []mastodon.ID{attachment.ID},
		Sensitive: false,
		//Visibility:  "unlisted",
		Visibility:  mastodon.VisibilityDirectMessage,
		Language:    "EN",
		ScheduledAt: &scheduledAt,
	}
	status, err := c.PostStatus(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
	//log.Info(fmt.Sprintf("Post scheduled at %s (ID: %s)", scheduledAt.String(), status.ID))
	log.Info(fmt.Sprintf("Toot sent (ID: %s)", status.ID))
}

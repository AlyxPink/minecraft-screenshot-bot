package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	launchGame()
	createNewWorld()
	setupScreenshot()
	for i := 0; i < SHOTS; i++ {
		log.Info(fmt.Sprintf("** Starting a new shot (%d/%d) **", i, SHOTS))
		setRandomTime()
		setRandomWeather()
		teleportPlayer()
		takeRandomScreenshot()
		path := getLatestScreenshot()
		postScreenshotToSocialMedia(path, i) // TODO: Run in background
	}
	quitGame()
	cleanup()
}

func cleanup() {
	// TODO: delete latest created world
	// TODO: delete screenshots files (?)
}

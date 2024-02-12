package main

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	screenshot_path := getLatestScreenshot()
	url := uploadToS3(screenshot_path)
	postScreenshotToSocialMedia(screenshot_path, url, 1)

	// launchGame()
	// createNewWorld()
	// setupScreenshot()
	// for i := 0; i < SHOTS; i++ {
	// 	log.Info(fmt.Sprintf("** Starting a new shot (%d/%d) **", i, SHOTS))
	// 	setRandomTime()
	// 	setRandomWeather()
	// 	teleportPlayer()
	// 	takeRandomScreenshot()
	//  screenshot := getLatestScreenshot()
	// 	postScreenshotToSocialMedia(screenshot, i)
	// }
	// quitGame()
	// cleanup()
}

func cleanup() {
	// TODO: delete latest created world
	// TODO: delete screenshots files (?)
}

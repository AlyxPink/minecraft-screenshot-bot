package main

import (
	"github.com/AlyxPink/minecraft-screenshot-bot/src"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	src.Start()
}

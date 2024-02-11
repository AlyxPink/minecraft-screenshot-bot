package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/charmbracelet/log"
	"github.com/sashabaranov/go-openai"
)

func describeImage(image_url url.URL) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Validate URL
	u, err := url.Parse(image_url.String())
	if err != nil || u.Scheme == "" || u.Host == "" {
		log.Fatal(fmt.Sprintf("Invalid URL %s: (%s)", image_url.String(), err))
	}

	respUrl, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: "Can you describe accurately the landscape of this Minecraft screenshot? Text must be under 1000 characters.",
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    u.String(),
							Detail: openai.ImageURLDetailLow,
						},
					},
				},
			},
		},
		MaxTokens: 250,
	})

	if err != nil {
		fmt.Printf("Error while describing image: %v\n", err)
		return "Randomly generated Minecraft screenshot by CraftViews bot."
	}
	return fmt.Sprint("Randomly generated Minecraft screenshot. AI generated alt-text:\" ", respUrl.Choices[0].Message.Content)
}

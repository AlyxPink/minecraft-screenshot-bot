package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/charmbracelet/log"
	"github.com/sashabaranov/go-openai"
)

func describeImage(screenshot_url url.URL) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	log.Info(fmt.Sprint("Sending request to OpenAI to get alt text..."))
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
							URL:    screenshot_url.String(),
							Detail: openai.ImageURLDetailLow,
						},
					},
				},
			},
		},
		MaxTokens: 250,
	})

	if err != nil {
		fmt.Printf("Error while describing image, returning default alt-text: %v\n", err)
		return "Randomly generated Minecraft screenshot by CraftViews bot."
	}

	altText := respUrl.Choices[0].Message.Content

	log.Info(fmt.Sprintf("Screenshot alt-text generated: %s", altText))

	return fmt.Sprint("Randomly generated Minecraft screenshot. AI generated alt-text:\" ", altText)
}

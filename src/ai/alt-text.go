package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/sashabaranov/go-openai"
)

type AltText struct {
	Long string
}

func DescribeImage(ctx context.Context, screenshot_url string) AltText {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	log.Info(fmt.Sprint("Sending request to OpenAI to get alt text..."))
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: "Can you describe accurately the landscape of this Minecraft screenshot? Text must be under 1000 characters but make it shorter if there isn't much to describe. Don't mention it's a Minecraft screenshot, we know that already.",
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    screenshot_url,
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
		return AltText{Long: "Minecraft screenshot by CraftViews bot."}
	}

	altText := resp.Choices[0].Message.Content

	log.Info(fmt.Sprintf("Screenshot alt-text generated: %s", altText))

	return AltText{Long: fmt.Sprintf("Minecraft screenshot. AI generated alt text: %s", altText)}
}

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

func DescribeImage(ctx context.Context, screenshot_url string) (altText AltText) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	log.FromContext(ctx).Info(fmt.Sprint("Sending request to OpenAI to get alt text..."))
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
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
		log.FromContext(ctx).Error("Error while describing image, returning default alt-text", "error", err)
		altText.Long = "Minecraft screenshot by CraftViews bot."
		return altText
	}

	altText.Long = fmt.Sprintf("Minecraft screenshot. AI generated alt text: %s", resp.Choices[0].Message.Content)
	log.FromContext(ctx).Info("Screenshot alt-text generated successfully", "alt-text", altText)

	return altText
}

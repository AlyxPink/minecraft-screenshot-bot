package describer

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct{}

func (describer OpenAI) GenerateFromURL(ctx context.Context, url string) (desc string, err error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	log.FromContext(ctx).Info(fmt.Sprint("Sending request to OpenAI to describe the image"))
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: PROMPT,
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    url,
							Detail: openai.ImageURLDetailLow,
						},
					},
				},
			},
		},
		MaxTokens: 250,
	})

	if err != nil {
		log.FromContext(ctx).Error("Error while describing image", "error", err)
		return "", err
	}

	desc = resp.Choices[0].Message.Content
	logSuccess(ctx, "OpenAI", desc)
	return desc, nil
}

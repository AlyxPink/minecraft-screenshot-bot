package describer

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
)

const (
	PROMPT = "Can you describe accurately the landscape of this Minecraft screenshot? Text must be under 1000 characters."
)

type Describer interface {
	GenerateFromURL(ctx context.Context, url string) (desc string, err error)
}

func Get() Describer {
	if os.Getenv("DEBUG") != "" {
		return &Debug{}
	}

	switch os.Getenv("DESCRIPTOR") {
	case "static":
		return &Static{}
	case "openai":
		return &OpenAI{}
	}

	return Static{}
}

func logSuccess(ctx context.Context, serviceName string, desc string) {
	log.FromContext(ctx).Info(
		"Image description generated successfully",
		"serviceName", serviceName,
		"description", desc,
	)
}

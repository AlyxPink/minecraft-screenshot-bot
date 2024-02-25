package describer

import (
	"context"

	"github.com/charmbracelet/log"
)

type Debug struct{}

func (ai Debug) GenerateFromURL(ctx context.Context, url string) (desc string, err error) {
	log.FromContext(ctx).Debug(
		"Debug describer called",
		"url", url,
	)

	return "", nil
}

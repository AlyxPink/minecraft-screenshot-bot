package describer

import (
	"context"
)

type Static struct{}

func (describer Static) GenerateFromURL(ctx context.Context, url string) (desc string, err error) {
	return "Minecraft screenshot made by CraftViews bot.", nil
}

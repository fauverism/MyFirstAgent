package anthropic

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
)

type Client struct {
	*anthropic.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client: anthropic.NewClient(apiKey),
	}
}

func (c *Client) SendMessage(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error) {
	message, err := c.Client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_7SonnetLatest,
		MaxTokens: int64(1024),
		Messages:  conversation,
	})
	return message, err
}
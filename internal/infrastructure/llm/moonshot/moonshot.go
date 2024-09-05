package moonshot

import (
	"context"
	"errors"
	"io"
	"shangxiehui-ai/config"
	"shangxiehui-ai/internal/common"
	"shangxiehui-ai/internal/facade/dto"

	"github.com/northes/go-moonshot"
)

type Client struct {
	moonshot *moonshot.Client
}

func (c *Client) parseRole(role string) moonshot.ChatCompletionsMessageRole {
	switch role {
	case common.ChatRoleSystem:
		return moonshot.RoleSystem
	case common.ChatRoleUser:
		return moonshot.RoleUser
	case common.ChatRoleAssistant:
		return moonshot.RoleAssistant
	default:
		return moonshot.RoleSystem
	}
}

func (c *Client) CompletionStream(ctx context.Context, messages []*dto.ChatMessage) (chan string, error) {

	requestMessages := make([]*moonshot.ChatCompletionsMessage, 0)

	for _, message := range messages {

		requestMessages = append(requestMessages, &moonshot.ChatCompletionsMessage{
			Role:    c.parseRole(message.Role),
			Content: message.Message,
		})
	}

	resp, err := c.moonshot.Chat().CompletionsStream(ctx, &moonshot.ChatCompletionsRequest{
		Model:       moonshot.ModelMoonshotV18K,
		Messages:    requestMessages,
		Temperature: 1,
		Stream:      true,
	})

	if err != nil {
		return nil, err
	}

	responseChan := make(chan string)

	go func() {
		for receive := range resp.Receive() {
			msg, err := receive.GetMessage()
			if err != nil {
				close(responseChan)
				if errors.Is(err, io.EOF) {
					break
				}
				break
			}
			if msg.Content != "" {
				responseChan <- msg.Content
			}
		}
	}()

	return responseChan, nil
}

func NewClient(cfg *config.Config) (*Client, error) {
	cli, err := moonshot.NewClientWithConfig(
		moonshot.NewConfig(
			moonshot.WithAPIKey(cfg.Moonshot.APIKey),
		),
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		moonshot: cli,
	}, nil
}

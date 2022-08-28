package nats

import (
	"chat/domain/nats"
	"fmt"
	"github.com/nats-io/nats.go"
)

type Client struct {
	cn *nats.Conn
}

// New initializes a connection to NATS server
func New() (*Client, error) {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to NATS: %v", err)
	}
	return &Client{cn: conn}, nil
}

func (c *Client) Sub(sub string) string {
	var out string
	c.cn.Subscribe(sub, func(msg *nats.Msg) {
		out = string(msg.Data)
	})
	return out
}

func (c *Client) Pub(form natsDomain.MessageForm) error {
	err := c.cn.Publish(form.Sub, []byte(form.Message))
	if err != nil {
		return err
	}
	return nil
}

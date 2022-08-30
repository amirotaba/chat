package PubSub

import (
	"chat/domain"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
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

func (c *Client) Sub(sub string) {
	c.cn.Subscribe(sub, func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})
}

func (c *Client) Pub(form domain.Message) {
	err := c.cn.Publish(form.Sub, []byte(form.Message))
	if err != nil {
		log.Println(err)
	}
}

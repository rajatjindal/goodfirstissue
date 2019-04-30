package slack

import (
	"github.com/bluele/slack"
)

//Tokens is for slack tokens
type Tokens struct {
	Token string `yaml:"token"`
}

//Client is slack client
type Client struct {
	slack *slack.Slack
}

//NewClient returns new twitter client
func NewClient(t *Tokens) (*Client, error) {
	return &Client{
		slack: slack.New(t.Token),
	}, nil
}

//JoinChannel joins channel
func (c *Client) JoinChannel(channel string) error {
	return c.slack.JoinChannel(channel)
}

//SendMessage sends msg to given channel
func (c *Client) SendMessage(channel, msg string) error {
	err := c.JoinChannel(channel)
	if err != nil {
		return err
	}

	return c.slack.ChatPostMessage(channel, msg, nil)
}

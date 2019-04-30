package slack

import (
	"io/ioutil"

	"github.com/bluele/slack"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var credentialsFile = "/var/openfaas/secrets/slack.yaml"

//Tokens is for slack tokens
type Tokens struct {
	Token string `yaml:"token"`
}

//Client is slack client
type Client struct {
	slack *slack.Slack
}

//NewClient returns new twitter client
func NewClient() (*Client, error) {
	r, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		logrus.Error("failed to read credentials", err.Error())
		return nil, err
	}

	t := &Tokens{}
	err = yaml.Unmarshal(r, t)
	if err != nil {
		logrus.Error("failed to unmarshal json", err.Error())
		return nil, err
	}

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

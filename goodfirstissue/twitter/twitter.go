package twitter

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"
)

var credentialsFile = "/var/openfaas/secrets/twitter.yaml"

//Tokens is for twitter tokens
type Tokens struct {
	ConsumerKey   string `yaml:"consumerKey"`
	ConsumerToken string `yaml:"consumerToken"`
	Token         string `yaml:"token"`
	TokenSecret   string `yaml:"tokenSecret"`
}

//Client is twitter client
type Client struct {
	twitter *twitter.Client
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

	config := oauth1.NewConfig(t.ConsumerKey, t.ConsumerToken)
	token := oauth1.NewToken(t.Token, t.TokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return &Client{
		twitter: twitter.NewClient(httpClient),
	}, nil
}

//Tweet tweets
func (c *Client) Tweet(msg string) {
	_, _, err := c.twitter.Statuses.Update(msg, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
}

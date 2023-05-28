package twitter

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	//lint:ignore SA1019 ignore this for now!
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
	"github.com/sirupsen/logrus"
)

const (
	maxTweetLength    = 280
	dotsAtTheEnd      = "...."
	newLinesCharCount = 5
)

// ErrTweetFailed is for failed tweet
var ErrTweetFailed = errors.New("failed to send tweet")

type RealClient struct {
	twitter *twitter.Client
}

var (
	_ socials.Provider = &RealClient{}
)

// NewClient returns new twitter client
func NewClient(client *http.Client, credsProvider creds.Provider) (*RealClient, error) {
	credentials, err := credsProvider.GetCredentials("twitter")
	if err != nil {
		return nil, err
	}

	config := oauth1.NewConfig(credentials["consumerKey"], credentials["consumerToken"])
	token := oauth1.NewToken(credentials["token"], credentials["tokenSecret"])
	httpClient := config.Client(context.WithValue(oauth1.NoContext, oauth1.HTTPClient, client), token)

	return &RealClient{
		twitter: twitter.NewClient(httpClient),
	}, nil
}

func (c *RealClient) Format(prefix string, event *github.IssuesEvent) string {
	fmt.Println("formatting the msg")
	return format(prefix, event)
}

func (c *RealClient) CreatePost(post string) error {
	fmt.Println("trying to post")
	_, _, err := c.twitter.Statuses.Update(post, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}
	fmt.Println("successful post")
	return nil
}

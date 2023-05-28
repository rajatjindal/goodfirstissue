package twitter

import (
	"errors"

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
func NewClient(credsProvider creds.Provider) (*RealClient, error) {
	credentials, err := credsProvider.GetCredentials("twitter")
	if err != nil {
		return nil, err
	}

	config := oauth1.NewConfig(credentials["consumerKey"], credentials["consumerToken"])
	token := oauth1.NewToken(credentials["token"], credentials["tokenSecret"])
	httpClient := config.Client(oauth1.NoContext, token)

	return &RealClient{
		twitter: twitter.NewClient(httpClient),
	}, nil
}

func (c *RealClient) Format(prefix string, event *github.IssuesEvent) string {
	return format(prefix, event)
}

func (c *RealClient) CreatePost(post string) error {
	_, _, err := c.twitter.Statuses.Update(post, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

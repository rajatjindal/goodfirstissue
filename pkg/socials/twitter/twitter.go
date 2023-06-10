package twitter

import (
	"context"
	"errors"
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

func (c *RealClient) CreatePost(ctx context.Context, prefix string, event *github.IssuesEvent) error {
	post := format(prefix, event)

	_, _, err := c.twitter.Statuses.Update(post, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (c *RealClient) Name() string {
	return "twitter"
}

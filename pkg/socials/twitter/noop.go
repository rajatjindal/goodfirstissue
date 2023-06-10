package twitter

import (
	"context"
	"fmt"

	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
)

type NoopClient struct{}

// NewClient returns new twitter client
func NewNoopClient(credsProvider creds.Provider) (*NoopClient, error) {
	return &NoopClient{}, nil
}

var (
	_ socials.Provider = &NoopClient{}
)

func (c *NoopClient) CreatePost(ctx context.Context, prefix string, event *github.IssuesEvent) error {
	fmt.Println(ctx, prefix, event)
	return nil
}

func (c *NoopClient) Name() string {
	return "noop"
}

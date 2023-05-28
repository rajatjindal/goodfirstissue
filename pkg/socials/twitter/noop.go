package twitter

import (
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

func (c *NoopClient) CreatePost(post string) error {
	fmt.Println(post)
	return nil
}

func (c *NoopClient) Format(prefix string, event *github.IssuesEvent) string {
	return format(prefix, event)
}

package socials

import (
	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/stretchr/testify/mock"
)

type FakeClient struct {
	mock.Mock
}

// NewClient returns new twitter client
func NewFakeClient(credsProvider creds.Provider) (*FakeClient, error) {
	return &FakeClient{}, nil
}

var (
	_ Provider = &FakeClient{}
)

func (c *FakeClient) CreatePost(post string) error {
	args := c.Called(post)
	return args.Error(0)
}

func (c *FakeClient) Format(prefix string, event *github.IssuesEvent) string {
	args := c.Called(prefix, event)
	return args.String(0)
}

package socials

import (
	"context"

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

func (c *FakeClient) Name() string {
	return "fake"
}

func (c *FakeClient) CreatePost(ctx context.Context, prefix string, event *github.IssuesEvent) error {
	args := c.Called(ctx, prefix, event)
	return args.Error(0)
}

package socials

import (
	"context"

	"github.com/google/go-github/v51/github"
)

type Provider interface {
	Name() string
	CreatePost(ctx context.Context, prefix string, event *github.IssuesEvent) error
}

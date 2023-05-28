package socials

import "github.com/google/go-github/v51/github"

type Provider interface {
	Format(prefix string, issue *github.Issue) string
	CreatePost(post string) error
}

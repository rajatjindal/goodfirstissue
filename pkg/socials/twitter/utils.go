package twitter

import (
	"fmt"

	"github.com/google/go-github/v51/github"
)

func format(prefix string, event *github.IssuesEvent) string {
	mention := getMention(event)

	prefixLength := len(prefix)
	mentionLength := len(mention)
	urlLength := len(event.Issue.GetHTMLURL())
	dotLength := len(dotsAtTheEnd)

	maxSummaryLength := maxTweetLength - (prefixLength + mentionLength + urlLength + dotLength + newLinesCharCount)
	title := event.Issue.GetTitle()
	if len(title) > maxSummaryLength {
		title = event.Issue.GetTitle()[0:maxSummaryLength] + dotsAtTheEnd
	}

	return fmt.Sprintf("%s %s\n\n%s\n%s", prefix, mention, title, event.Issue.GetHTMLURL())
}

func getMention(event *github.IssuesEvent) string {
	tHandle := ""
	ok := false
	if tHandle, ok = twitterMap[event.Issue.Repository.GetFullName()]; ok && tHandle != "" {
		return fmt.Sprintf(" @%s", tHandle)
	}

	if tHandle, ok = twitterMap[event.Issue.Repository.Owner.GetLogin()]; ok && tHandle != "" {
		return fmt.Sprintf(" @%s", tHandle)
	}

	return ""
}

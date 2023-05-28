package twitter

import (
	"fmt"

	"github.com/google/go-github/v51/github"
)

func format(prefix string, issue *github.Issue) string {
	mention := getMention(issue)

	prefixLength := len(prefix)
	mentionLength := len(mention)
	urlLength := len(issue.GetHTMLURL())
	dotLength := len(dotsAtTheEnd)

	maxSummaryLength := maxTweetLength - (prefixLength + mentionLength + urlLength + dotLength + newLinesCharCount)
	title := issue.GetTitle()
	if len(title) > maxSummaryLength {
		title = issue.GetTitle()[0:maxSummaryLength] + dotsAtTheEnd
	}

	return fmt.Sprintf("%s %s\n\n%s\n%s", prefix, mention, title, issue.GetHTMLURL())
}

func getMention(issue *github.Issue) string {
	tHandle := ""
	ok := false
	if tHandle, ok = twitterMap[issue.Repository.GetFullName()]; ok && tHandle != "" {
		return fmt.Sprintf(" @%s", tHandle)
	}

	if tHandle, ok = twitterMap[issue.Repository.Owner.GetLogin()]; ok && tHandle != "" {
		return fmt.Sprintf(" @%s", tHandle)
	}

	return ""
}

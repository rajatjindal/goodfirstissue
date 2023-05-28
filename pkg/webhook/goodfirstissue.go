package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v51/github"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoActionRequired = errors.New("no action required")
	ErrUnsupportedEvent = errors.New("event not supported")
)

func getPrefix(event *github.IssuesEvent) string {
	msg := ""
	switch event.GetAction() {
	case "opened":
		msg = "a new #goodfirstissue opened"
	case "reopened":
		msg = "a #goodfirstissue got reopened"
	case "labeled":
		if event.Issue.GetState() == "open" {
			msg = "an issue just got labeled #goodfirstissue"
		}
	case "unassigned":
		if event.Issue.GetState() == "open" {
			msg = "an issue just got available for assignment #goodfirstissue"
		}
	default:
		logrus.Warnf("unsupported event action %s", event.GetAction())
		return ""
	}

	msg += fmt.Sprintf(" for %s", event.Repo.GetFullName())
	if event.Repo.GetFullName() != "codinasion/program" && event.Repo.Language != nil {
		msg += fmt.Sprintf(" #%s", event.Repo.GetLanguage())
	}

	return msg
}

func isGoodFirstIssue(msgType string, payload []byte) (*github.IssuesEvent, error) {
	event, err := parseEvent(msgType, payload)
	if err != nil {
		return nil, err
	}

	if event == nil || event.Issue == nil {
		return nil, fmt.Errorf("event is malformed")
	}

	if !goodFirstIssue(event.Issue.Labels) {
		return nil, ErrNoActionRequired
	}

	return event, nil
}

func parseEvent(msgType string, payload []byte) (*github.IssuesEvent, error) {
	x := "IssuesEvent"
	rawEvent := github.Event{
		Type:       &x,
		RawPayload: (*json.RawMessage)(&payload),
	}

	raw, err := rawEvent.ParsePayload()
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload. Error: %v", err)
	}

	event, ok := raw.(*github.IssuesEvent)
	if !ok {
		return nil, ErrUnsupportedEvent
	}

	return event, nil
}

func goodFirstIssue(labels []*github.Label) bool {
	for _, l := range labels {
		labelName := strings.ToLower(l.GetName())
		if labelName == "good first issue" || labelName == "good-first-issue" {
			return true
		}

		if strings.Contains(labelName, "good") &&
			strings.Contains(labelName, "first") &&
			strings.Contains(labelName, "issue") {
			return true
		}
	}

	return false
}

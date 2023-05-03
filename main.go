package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/cache"
	"github.com/rajatjindal/goodfirstissue/logrus"
	"github.com/rajatjindal/goodfirstissue/twitter"
)

const (
	credentialsFile   = "/etc/secrets/goodfirstissue.yaml"
	cacheExpiration   = 1 * time.Minute
	maxTweetLength    = 280
	dotsAtTheEnd      = "...."
	newLinesCharCount = 5
)

type WebhookHandler struct {
	Twitter          *twitter.Client
	TwitterHandleMap map[string]string
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		tokens := &twitter.Tokens{
			ConsumerKey:   "",
			ConsumerToken: "",
			Token:         "",
			TokenSecret:   "",
		}

		spinhttpclient := spinhttp.NewClient()
		client, err := twitter.NewClient(spinhttpclient, tokens)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		twitterHandleMap := twitter.GetTwitterHandleMap()
		handler := &WebhookHandler{
			Twitter:          client,
			TwitterHandleMap: twitterHandleMap,
		}

		if r.URL.Path == "/webhook" {
			handler.Handle(w, r)
			return
		}

		if r.URL.Path == "/cleanup-cache" {
			handler.CleanCache(w, r)
			return
		}

		fmt.Fprintln(w, "hello from goodfirstissue running on Fermyon Cloud")
	})
}

func main() {}

// Handle handles the function call to function
func (h *WebhookHandler) CleanCache(w http.ResponseWriter, r *http.Request) {
	err := cache.CleanupCache()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handle handles the function call to function
func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventForSpin(r)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "failed to process the event", http.StatusInternalServerError)
		return
	}

	// not a good first issue
	if !goodFirstIssue(event.Issue.Labels) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return
	}

	// already tweeted about it few mins back
	if _, found := cache.Get(event.Issue.GetHTMLURL()); found {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return
	}

	// format the tweet
	msg := h.getMsg(event)
	if msg == "" {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return
	}

	err = h.Twitter.Tweet(msg)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cache.Set(event.Issue.GetHTMLURL(), true)
	if err != nil {
		logrus.Warnf("failed to cache. Error %v", err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(msg))
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

func parseEvent(r *http.Request) (*github.IssuesEvent, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	t := github.WebHookType(r)
	if t == "" {
		return nil, fmt.Errorf("header 'X-GitHub-Event' not found")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body. Error: %v", err)
	}

	raw, err := github.ParseWebHook(t, body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload. Error: %v", err)
	}

	event, ok := raw.(*github.IssuesEvent)
	if !ok {
		return nil, fmt.Errorf("event not supported. Error: %v", err)
	}

	return event, nil
}

func parseEventForSpin(r *http.Request) (*github.IssuesEvent, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	webhookType := github.WebHookType(r)
	if webhookType == "" {
		return nil, fmt.Errorf("header 'X-GitHub-Event' not found")
	}

	if webhookType != "issues" {
		return nil, fmt.Errorf("unsupported event %s", webhookType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body. Error: %v", err)
	}

	x := "IssuesEvent"
	rawEvent := github.Event{
		Type:       &x,
		RawPayload: (*json.RawMessage)(&body),
	}

	raw, err := rawEvent.ParsePayload()
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload. Error: %v", err)
	}

	event, ok := raw.(*github.IssuesEvent)
	if !ok {
		return nil, fmt.Errorf("event not supported. type is %T", raw)
	}

	return event, nil
}

func (h *WebhookHandler) getMsg(event *github.IssuesEvent) string {
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

	//send to twitter
	msg += fmt.Sprintf(" for %s", event.Repo.GetFullName())
	if event.Repo.GetFullName() != "codinasion/program" && event.Repo.Language != nil {
		msg += fmt.Sprintf(" #%s", event.Repo.GetLanguage())
	}

	//if we have entry for specific repo use that,
	//else fallback to check if entry exist for owner of repo
	tHandle := ""
	ok := false
	if tHandle, ok = h.TwitterHandleMap[event.Repo.GetFullName()]; ok && tHandle != "" {
		msg += fmt.Sprintf(" @%s", tHandle)
	}

	if tHandle == "" {
		if tHandle, ok = h.TwitterHandleMap[event.Repo.Owner.GetLogin()]; ok && tHandle != "" {
			msg += fmt.Sprintf(" @%s", tHandle)
		}
	}

	msgLength := len(msg)
	urlLength := len(event.Issue.GetHTMLURL())
	dotLength := len(dotsAtTheEnd)

	maxSummaryLength := maxTweetLength - (msgLength + urlLength + dotLength + newLinesCharCount)
	summary := event.Issue.GetTitle()

	if len(summary) > maxSummaryLength {
		summary = summary[0:maxSummaryLength] + dotsAtTheEnd
	}

	msg += fmt.Sprintf("\n\n%s\n%s", summary, event.Issue.GetHTMLURL())

	return msg
}

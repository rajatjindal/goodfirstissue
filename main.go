package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rajatjindal/goodfirstissue/twitter"

	"github.com/google/go-github/v51/github"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
	cache            *gocache.Cache
}

func main() {
	secrets, err := initCredentials()
	if err != nil {
		logrus.Fatalf("failed to init creds %v", err)
	}

	client, err := twitter.NewClient(secrets)
	if err != nil {
		logrus.Fatalf("failed to create twitter client %v", err)
	}

	twitterHandleMap := twitter.GetTwitterHandleMap()
	handler := &WebhookHandler{
		Twitter:          client,
		TwitterHandleMap: twitterHandleMap,
		cache:            gocache.New(cacheExpiration, 2*time.Minute),
	}

	http.HandleFunc("/", handler.Handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func initCredentials() (*twitter.Tokens, error) {
	r, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file with err: %s", err.Error())
	}

	t := &twitter.Tokens{}
	err = yaml.Unmarshal(r, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Handle handles the function call to function
func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
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
	if _, found := h.cache.Get(event.Issue.GetHTMLURL()); found {
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
	logrus.Tracef("%s", string(body))

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

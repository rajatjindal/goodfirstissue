package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rajatjindal/goodfirstissue/twitter"

	"github.com/google/go-github/github"
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
	http.ListenAndServe(":8080", nil)
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
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body cannot be empty"))
		return
	}

	t := github.WebHookType(r)
	if t == "" {
		logrus.Error("header 'X-GitHub-Event' not found. cannot handle this request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("header 'X-GitHub-Event' not found."))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("failed to read request body. error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to read request body."))
		return
	}
	logrus.Tracef("%s", string(body))

	e, err := github.ParseWebHook(t, body)
	if err != nil {
		logrus.Error("failed to parsepayload. error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to parse payload."))
		return
	}

	msg := ""
	if o, ok := e.(*github.IssuesEvent); ok && goodFirstIssue(o.Issue.Labels) {
		if _, found := h.cache.Get(stringValue(o.Issue.HTMLURL)); found {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}

		//TODO: add logic to not retweet same issue
		switch stringValue(o.Action) {
		case "opened":
			msg = "a new #goodfirstissue opened"
		case "reopened":
			msg = "a #goodfirstissue got reopened"
		case "labeled":
			if stringValue(o.Issue.State) == "open" {
				msg = "an issue just got labeled #goodfirstissue"
			}
		case "unassigned":
			if stringValue(o.Issue.State) == "open" {
				msg = "an issue just got available for assignment #goodfirstissue"
			}
		}

		//send to twitter
		if msg != "" {
			msg += fmt.Sprintf(" for %s", stringValue(o.Repo.FullName))
			if stringValue(o.Repo.FullName) != "codinasion/program" && o.Repo.Language != nil {
				msg += fmt.Sprintf(" #%s", stringValue(o.Repo.Language))
			}

			//if we have entry for specific repo use that,
			//else fallback to check if entry exist for owner of repo
			tHandle := ""
			ok := false
			if tHandle, ok = h.TwitterHandleMap[stringValue(o.Repo.FullName)]; ok && tHandle != "" {
				msg += fmt.Sprintf(" @%s", tHandle)
			}

			if tHandle == "" {
				if tHandle, ok = h.TwitterHandleMap[stringValue(o.Repo.Owner.Login)]; ok && tHandle != "" {
					msg += fmt.Sprintf(" @%s", tHandle)
				}
			}

			msgLength := len(msg)
			urlLength := len(o.Issue.GetHTMLURL())
			dotLength := len(dotsAtTheEnd)

			maxSummaryLength := maxTweetLength - (msgLength + urlLength + dotLength + newLinesCharCount)
			summary := o.Issue.GetTitle()

			if len(summary) > maxSummaryLength {
				summary = summary[0:maxSummaryLength] + dotsAtTheEnd
			}

			msg += fmt.Sprintf("\n\n%s\n%s", summary, o.Issue.GetHTMLURL())
			h.cache.Set(stringValue(o.Issue.HTMLURL), "tweeted", cacheExpiration)
			h.Twitter.Tweet(msg)
		}
	}

	if msg == "" {
		msg = "OK"
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func goodFirstIssue(labels []github.Label) bool {
	for _, l := range labels {
		if stringValue(l.Name) == "good first issue" || stringValue(l.Name) == "good-first-issue" {
			return true
		}

		if strings.Contains(stringValue(l.Name), "good") &&
			strings.Contains(stringValue(l.Name), "first") &&
			strings.Contains(stringValue(l.Name), "issue") {
			return true
		}
	}

	return false
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

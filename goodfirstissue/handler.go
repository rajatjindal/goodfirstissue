package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"handler/function/twitter"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var (
	twitterClient        *twitter.Client
	twitterClientInitErr error
)

var twitterMap = map[string]string{}

func init() {
	twitterClient, twitterClientInitErr = twitter.NewClient()

	f := "twitter-handle-map.yaml"
	c, err := ioutil.ReadFile(f)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}

	err = yaml.Unmarshal(c, &twitterMap)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
}

//Handle handles the function call
func Handle(w http.ResponseWriter, r *http.Request) {
	if twitterClient == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("twitter client not initialized. error: %s", twitterClientInitErr.Error())))
		return
	}

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

	body, err := ioutil.ReadAll(r.Body)
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

		if msg != "" {
			msg += fmt.Sprintf(" for %s.\n\n%s", stringValue(o.Repo.FullName), stringValue(o.Issue.HTMLURL))
			if o.Repo.Language != nil {
				msg += fmt.Sprintf("#%s", stringValue(o.Repo.Language))
			}

			//if we have entry for specific repo use that,
			//else fallback to check if entry exist for owner of repo
			tHandle := ""
			ok := false
			if tHandle, ok = twitterMap[stringValue(o.Repo.FullName)]; ok && tHandle != "" {
				msg += fmt.Sprintf(" @%s", tHandle)
			}

			if tHandle == "" {
				if tHandle, ok = twitterMap[stringValue(o.Repo.Owner.Login)]; ok && tHandle != "" {
					msg += fmt.Sprintf(" @%s", tHandle)
				}
			}

			twitterClient.Tweet(msg)
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

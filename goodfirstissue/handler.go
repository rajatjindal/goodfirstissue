package function

// package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"handler/function/twitter"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

var (
	twitterClient        *twitter.Client
	twitterClientInitErr error
)

func init() {
	twitterClient, twitterClientInitErr = twitter.NewClient()
}

// func main() {
// 	http.HandleFunc("/", Handle)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}
// }

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

	//per https://developer.github.com/webhooks/#events integration_installation and integration_installation_repositories events are
	// deprecated and replaced by installation and installation_repositories. But github is still sending integration_installation and
	// integration_installation_repositories and go-github don't understand these.

	switch t {
	case "integration_installation":
		t = "installation"
	case "integration_installation_repositories":
		t = "installation_repositories"
	}

	e, err := github.ParseWebHook(t, body)
	if err != nil {
		logrus.Error("failed to parsepayload. error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to parse payload."))
		return
	}

	var msgs []string

	switch o := e.(type) {
	case *github.IssuesEvent:
		msgs = handleIssuesEvent(o)
	case *github.InstallationEvent:
		msgs = handleInstallationEvent(o)
	case *github.InstallationRepositoriesEvent:
		msgs = handleInstallationRepositoriesEvent(o)
	}

	for _, msg := range msgs {
		twitterClient.Tweet(msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handleIssuesEvent(o *github.IssuesEvent) []string {
	if !goodFirstIssue(o.Issue.Labels) {
		return nil
	}

	msg := ""
	switch {
	case *o.Action == "opened":
		msg = "a new #goodfirstissue opened"
	case *o.Action == "reopened":
		msg = "a #goodfirstissue got reopened"
	case *o.Action == "labeled":
		msg = "an issue just got labeled #goodfirstissue"
	case *o.Action == "unassigned":
		msg = "an issue just got available for assignment #goodfirstissue"
	}

	if msg != "" {
		msg += fmt.Sprintf(" for %s.\n#%s\n%s", *o.Repo.FullName, *o.Repo.Language, *o.Issue.HTMLURL)
	}

	return []string{msg}
}

func goodFirstIssue(labels []github.Label) bool {
	for _, l := range labels {
		if *l.Name == "good first issue" || *l.Name == "good-first-issue" {
			return true
		}

		if strings.Contains(*l.Name, "good") &&
			strings.Contains(*l.Name, "first") &&
			strings.Contains(*l.Name, "issue") {
			return true
		}
	}

	return false
}

func handleInstallationEvent(o *github.InstallationEvent) []string {
	if stringValue(o.Action) != "created" {
		return nil
	}

	msgs := []string{}

	for _, r := range o.Repositories {
		htmlURL := fmt.Sprintf("%s/%s", stringValue(o.Installation.Account.HTMLURL), stringValue(r.Name))
		msgs = append(msgs, fmt.Sprintf("yay!! Lets welcome %s to #goodfirstissue", htmlURL))
	}

	return msgs
}

func handleInstallationRepositoriesEvent(o *github.InstallationRepositoriesEvent) []string {
	if stringValue(o.Action) != "added" {
		return nil
	}

	msgs := []string{}
	for _, r := range o.RepositoriesAdded {
		htmlURL := fmt.Sprintf("%s/%s", stringValue(o.Installation.Account.HTMLURL), stringValue(r.Name))
		msgs = append(msgs, fmt.Sprintf("yay!! Lets welcome %s to #goodfirstissue", htmlURL))
	}

	return msgs
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

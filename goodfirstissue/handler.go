package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

//Handle handles the function call
func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body cannot be empty"))
	}

	t := github.WebHookType(r)
	if t == "" {
		logrus.Error("header 'X-GitHub-Event' not found. cannot handle this request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("header 'X-GitHub-Event' not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error("failed to read request body. error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	logrus.Tracef("%s", string(body))

	e, err := github.ParseWebHook(t, body)
	if err != nil {
		logrus.Error("failed to parsepayload. error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	msg := ""
	if o, ok := e.(*github.IssuesEvent); ok && goodFirstIssue(o.Issue.Labels) {
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
			msg += fmt.Sprintf(" for %s\n%s", *o.Repo.FullName, *o.Issue.HTMLURL)
		}
	}

	if msg == "" {
		//comment 1
		msg = "OK"
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
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

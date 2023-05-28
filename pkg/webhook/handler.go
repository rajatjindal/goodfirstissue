package webhook

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/cache"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	socialProvider socials.Provider
	cacheProvider  cache.Provider
}

func NewHandler(cacheProvider cache.Provider, socialProvider socials.Provider) *Handler {
	return &Handler{
		socialProvider: socialProvider,
		cacheProvider:  cacheProvider,
	}
}

// Handle handles the function call to function
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	code, body := h.handle(r)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func (h *Handler) handle(r *http.Request) (int, []byte) {
	fmt.Println("entering handle function")
	if r.Body == nil {
		logrus.Error("request body is missing")
		return http.StatusBadRequest, []byte("bad request")
	}

	fmt.Println("getting webhook type")
	msgType := github.WebHookType(r)
	if msgType == "" {
		logrus.Error("github event header is missing")
		return http.StatusBadRequest, []byte("bad request")
	}

	fmt.Println("reading payload")
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("failed to read request body. Error: %v", err)
		return http.StatusInternalServerError, []byte("failed to read request body")
	}

	fmt.Println("calling isGoodFirstIssue function")
	event, err := isGoodFirstIssue(msgType, payload)
	if errors.Is(err, ErrNoActionRequired) {
		return http.StatusOK, []byte("OK")
	}

	fmt.Println("checking unsupportedevent error")
	if errors.Is(err, ErrUnsupportedEvent) {
		return http.StatusBadRequest, []byte("bad request")
	}

	fmt.Println("checking if err is nil")
	if err != nil {
		logrus.Error(err.Error())
		return http.StatusInternalServerError, []byte(err.Error())
	}

	fmt.Println("checking in cache")
	if _, ok := h.cacheProvider.Get(event.Issue.GetHTMLURL()); ok {
		return http.StatusOK, []byte("OK")
	}

	fmt.Println("getting prefix")
	prefix := getPrefix(event)

	fmt.Println("posting to twitter")
	err = h.socialProvider.CreatePost(h.socialProvider.Format(prefix, event))
	if err != nil {
		logrus.Error(err.Error())
		return http.StatusInternalServerError, []byte(err.Error())
	}

	fmt.Println("setting in cache")
	err = h.cacheProvider.Set(event.Issue.GetHTMLURL(), true)
	if err != nil {
		logrus.Warn(err)
	}

	return http.StatusOK, []byte("OK")
}

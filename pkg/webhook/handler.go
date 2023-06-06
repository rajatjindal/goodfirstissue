package webhook

import (
	"errors"
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
	if r.Body == nil {
		logrus.Error("request body is missing")
		return http.StatusBadRequest, []byte("bad request")
	}

	msgType := github.WebHookType(r)
	if msgType == "" {
		logrus.Error("github event header is missing")
		return http.StatusBadRequest, []byte("bad request")
	}

	if msgType != "issues" {
		logrus.Errorf("unsupported github event %q", msgType)
		return http.StatusBadRequest, []byte("bad request")
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("failed to read request body. Error: %v", err)
		return http.StatusInternalServerError, []byte("failed to read request body")
	}

	event, err := isGoodFirstIssue(msgType, payload)
	if errors.Is(err, ErrNoActionRequired) {
		return http.StatusOK, []byte("OK")
	}

	if errors.Is(err, ErrUnsupportedEvent) {
		return http.StatusBadRequest, []byte("bad request")
	}

	if err != nil {
		logrus.Error(err.Error())
		return http.StatusInternalServerError, []byte(err.Error())
	}

	if _, ok := h.cacheProvider.Get(event.Issue.GetHTMLURL()); ok {
		return http.StatusOK, []byte("OK")
	}

	prefix, ok := getPrefix(event)
	if !ok {
		return http.StatusOK, []byte("OK")
	}

	err = h.socialProvider.CreatePost(h.socialProvider.Format(prefix, event))
	if err != nil {
		logrus.Error(err.Error())
		return http.StatusInternalServerError, []byte(err.Error())
	}

	err = h.cacheProvider.Set(event.Issue.GetHTMLURL(), true)
	if err != nil {
		logrus.Warn(err)
	}

	return http.StatusOK, []byte("OK")
}

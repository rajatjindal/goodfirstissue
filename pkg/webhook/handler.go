package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/cache"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	socialProviders []socials.Provider
	cacheProvider   cache.Provider
}

func NewHandler(cacheProvider cache.Provider, socialProviders ...socials.Provider) *Handler {
	return &Handler{
		socialProviders: socialProviders,
		cacheProvider:   cacheProvider,
	}
}

// Handle handles the function call to function
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	code, body := h.handle(r)
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func (h *Handler) handle(r *http.Request) (int, []byte) {
	stats := map[string]time.Duration{}
	startTime := time.Now()
	defer func() {
		stats["handler"] = time.Since(startTime)
		d, _ := json.Marshal(stats)
		fmt.Println(string(d))
	}()

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

	var errs []error
	for _, provider := range h.socialProviders {
		stime := time.Now()
		err = provider.CreatePost(r.Context(), prefix, event)
		if err != nil {
			logrus.Error(err.Error())
			errs = append(errs, fmt.Errorf("%s: %v", provider.Name(), err))
		}
		stats[fmt.Sprintf("post:%s", provider.Name())] = time.Since(stime)
	}

	if len(errs) > 0 {
		msg := ""
		for _, e := range errs {
			msg += e.Error() + "\n"
		}

		return http.StatusInternalServerError, []byte(msg)
	}

	err = h.cacheProvider.Set(event.Issue.GetHTMLURL(), true)
	if err != nil {
		logrus.Warn(err)
	}

	return http.StatusOK, []byte("OK")
}

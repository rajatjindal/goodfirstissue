package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-github/v51/github"
	"github.com/rajatjindal/goodfirstissue/pkg/cache"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler(t *testing.T) {
	testcases := []struct {
		name         string
		setupRequest func() *http.Request
		setupHandler func(c *cache.FakeCache, s *socials.FakeClient) *Handler
		expectedCode int
		expectedBody []byte
	}{
		{
			name: "body is nil",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", nil)
				return req
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: []byte("bad request"),
		},
		{
			name: "github header is missing",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				return req
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: []byte("bad request"),
		},
		{
			name: "not an issues event",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				req.Header.Set("X-Github-Event", "pull_request")
				return req
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: []byte("bad request"),
		},
		{
			name: "action opened, not goodfirstissue",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				event.Issue.Labels = []*github.Label{}
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				req.Header.Set("X-Github-Event", "issues")
				return req
			},
			expectedCode: http.StatusOK,
			expectedBody: []byte("OK"),
		},
		{
			name: "action opened, goodfirstissue",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				c.On("Get", "https://github.com/rajatjindal/goodfirstissue/issues/36").Return(false, false)
				c.On("Set", "https://github.com/rajatjindal/goodfirstissue/issues/36", true).Return(nil)
				s.On("Format", mock.AnythingOfType("string"), mock.AnythingOfType("*github.Issue")).Return("formatted msg")
				s.On("CreatePost", "formatted msg").Return(nil)
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				req.Header.Set("X-Github-Event", "issues")
				return req
			},
			expectedCode: http.StatusOK,
			expectedBody: []byte("OK"),
		},
		{
			name: "action opened, goodfirstissue, cached",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				c.On("Get", "https://github.com/rajatjindal/goodfirstissue/issues/36").Return(true, true)
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				req.Header.Set("X-Github-Event", "issues")
				return req
			},
			expectedCode: http.StatusOK,
			expectedBody: []byte("OK"),
		},
		{
			name: "action opened, goodfirstissue, create post failed",
			setupHandler: func(c *cache.FakeCache, s *socials.FakeClient) *Handler {
				c.On("Get", "https://github.com/rajatjindal/goodfirstissue/issues/36").Return(false, false)
				s.On("Format", mock.AnythingOfType("string"), mock.AnythingOfType("*github.Issue")).Return("formatted msg")
				s.On("CreatePost", "formatted msg").Return(fmt.Errorf("failed to create post"))
				return NewHandler(c, s)
			},
			setupRequest: func() *http.Request {
				event := testEvent()
				req, _ := http.NewRequest(http.MethodPost, "https://example.com", reader(event))
				req.Header.Set("X-Github-Event", "issues")
				return req
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: []byte("failed to create post"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			fakeCache := cache.NewFakeCache(1*time.Minute, 2*time.Minute)
			fakeProvider := &socials.FakeClient{}

			handler := tc.setupHandler(fakeCache, fakeProvider)
			req := tc.setupRequest()

			code, body := handler.handle(req)
			assert.Equal(t, tc.expectedCode, code)
			assert.Equal(t, tc.expectedBody, body)

			fakeCache.AssertExpectations(t)
			fakeProvider.AssertExpectations(t)
		})
	}
}

func testEvent() *github.IssuesEvent {
	return &github.IssuesEvent{
		Action: ptr("opened"),
		Issue: &github.Issue{
			HTMLURL: ptr("https://github.com/rajatjindal/goodfirstissue/issues/36"),
			Repository: &github.Repository{
				Name:     ptr("goodfirstissue"),
				FullName: ptr("rajatjindal/goodfirstissue"),
				Language: ptr("golang"),
			},
			Labels: []*github.Label{
				{
					Name: ptr("good first issue"),
				},
			},
		},
	}
}

func reader(g *github.IssuesEvent) io.Reader {
	b, _ := json.Marshal(g)
	return bytes.NewReader(b)
}

func ptr[T any](v T) *T {
	return &v
}

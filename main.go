package main

import (
	"net/http"
	"time"

	"github.com/rajatjindal/goodfirstissue/pkg/cache"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials/twitter"
	"github.com/rajatjindal/goodfirstissue/pkg/webhook"

	"github.com/sirupsen/logrus"
)

const (
	credentialsFile = "/etc/secrets/goodfirstissue.yaml"
)

func main() {
	credsProvider := creds.FileProvider(credentialsFile)
	twitter, err := twitter.NewClient(credsProvider)
	if err != nil {
		logrus.Fatal(err)
	}

	cacheProvider := cache.NewGoCache(1*time.Minute, 2*time.Minute)
	handler := webhook.NewHandler(cacheProvider, twitter)

	http.HandleFunc("/", handler.Handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

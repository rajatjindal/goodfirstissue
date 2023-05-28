package main

import (
	"fmt"
	"net/http"
	"time"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/rajatjindal/goodfirstissue/pkg/cache"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials/twitter"
	"github.com/rajatjindal/goodfirstissue/pkg/webhook"

	"github.com/sirupsen/logrus"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entering spin Handle func")
		credsProvider := creds.KVProvider()
		client := spinhttp.NewClient()

		twitter, err := twitter.NewClient(client, credsProvider)
		if err != nil {
			logrus.Fatal(err)
		}

		cacheProvider := cache.KVProvider(1*time.Minute, 2*time.Minute)
		handler := webhook.NewHandler(cacheProvider, twitter)

		if r.URL.Path == "/cleanup-cache" {
			err := cacheProvider.CleanupExpiredCache()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}

		handler.Handle(w, r)
	})
}

func main() {}

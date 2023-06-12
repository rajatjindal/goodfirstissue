package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/rajatjindal/goodfirstissue/pkg/cache/kvcache"
	"github.com/rajatjindal/goodfirstissue/pkg/creds/kvcreds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials/bluesky"
	"github.com/rajatjindal/goodfirstissue/pkg/socials/twitter"
	"github.com/rajatjindal/goodfirstissue/pkg/webhook"

	"github.com/sirupsen/logrus"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		stime := time.Now()
		stats := map[string]interface{}{}
		defer func() {
			stats["total"] = time.Since(stime)
			stats["guid"] = r.Header.Get("X-GitHub-Delivery")
			d, _ := json.Marshal(stats)
			fmt.Println(string(d))
		}()

		fmt.Println("entering spin Handle func")
		credsProvider := kvcreds.Provider()
		client := spinhttp.NewClient()

		twitterTime := time.Now()
		twitter, err := twitter.NewClient(client, credsProvider)
		if err != nil {
			logrus.Fatal(err)
		}
		stats["creds:twitter"] = time.Since(twitterTime)

		bskyTime := time.Now()
		bluesky, err := bluesky.NewClient(client, credsProvider)
		if err != nil {
			logrus.Fatal(err)
		}
		stats["creds:bluesky"] = time.Since(bskyTime)

		cacheProvider := kvcache.Provider(1*time.Minute, 2*time.Minute)
		handler := webhook.NewHandler(cacheProvider, twitter, bluesky)

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

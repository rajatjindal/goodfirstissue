package main

import (
	"net/http"
	"time"

	"github.com/rajatjindal/goodfirstissue/pkg/cache/gocache"
	"github.com/rajatjindal/goodfirstissue/pkg/creds/filecreds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials/twitter"
	"github.com/rajatjindal/goodfirstissue/pkg/webhook"

	"github.com/sirupsen/logrus"
)

// func init() {
// 	spinhttp.Handle(Entrypoint)
// }

func main() {
	credsProvider := filecreds.Provider("/etc/secrets/goodfirstissue.yaml")

	twitter, err := twitter.NewClient(http.DefaultClient, credsProvider)
	if err != nil {
		logrus.Fatal(err)
	}

	cacheProvider := gocache.Provider(1*time.Minute, 2*time.Minute)
	handler := webhook.NewHandler(cacheProvider, twitter)

	http.HandleFunc("/cleanup-cache", func(w http.ResponseWriter, r *http.Request) {
		err := cacheProvider.CleanupExpiredCache()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", handler.Handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

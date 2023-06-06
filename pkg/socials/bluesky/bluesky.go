package bluesky

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/google/go-github/v51/github"
	"github.com/otiai10/opengraph"
	"github.com/rajatjindal/goodfirstissue/pkg/creds"
	"github.com/rajatjindal/goodfirstissue/pkg/socials"
	"github.com/sirupsen/logrus"
)

const (
	base = "https://bsky.social"
)

type BlueSky struct {
	xrpcc *xrpc.Client
}

var (
	_ socials.Provider = &BlueSky{}
)

func NewClient(client *http.Client, credsProvider creds.Provider) (*BlueSky, error) {
	credentials, err := credsProvider.GetCredentials("bluesky")
	if err != nil {
		return nil, err
	}

	xrpcc := &xrpc.Client{
		Auth: &xrpc.AuthInfo{
			Handle: credentials["handle"],
		},
		Client: client,
		Host:   base,
	}

	auth, err := comatproto.ServerCreateSession(context.TODO(), xrpcc, &comatproto.ServerCreateSession_Input{
		Identifier: xrpcc.Auth.Handle,
		Password:   credentials["password"],
	})
	if err != nil {
		return nil, err
	}

	xrpcc.Auth.AccessJwt = auth.AccessJwt
	xrpcc.Auth.RefreshJwt = auth.RefreshJwt
	xrpcc.Auth.Did = auth.Did
	xrpcc.Auth.Handle = auth.Handle

	return &BlueSky{
		xrpcc: xrpcc,
	}, nil
}

func (b *BlueSky) CreatePost(ctx context.Context, prefix string, event *github.IssuesEvent) error {
	post, err := b.format(ctx, prefix, event)
	if err != nil {
		return err
	}

	_, err = comatproto.RepoCreateRecord(ctx, b.xrpcc, post)
	return err
}

func getImage(ctx context.Context, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image")
	}

	return io.ReadAll(resp.Body)
}

func (b *BlueSky) format(ctx context.Context, prefix string, event *github.IssuesEvent) (*comatproto.RepoCreateRecord_Input, error) {
	postMsg := prefix + " \n " + event.Issue.GetTitle()
	post := &appbsky.FeedPost{
		Text:      postMsg,
		CreatedAt: time.Now().Format(time.RFC3339),
		Facets:    DetectFacets(postMsg),
	}

	og, err := b.getEmbedData(ctx, event)
	if err != nil {
		logrus.Warnf("failed to fetch embed data. error: %v", err)
	}

	if og != nil {
		post.Embed = og
	}

	return &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       b.xrpcc.Auth.Did,
		Record: &util.LexiconTypeDecoder{
			Val: post,
		},
	}, nil
}

func (b *BlueSky) Name() string {
	return "bluesky"
}

func (b *BlueSky) getEmbedData(ctx context.Context, event *github.IssuesEvent) (*appbsky.FeedPost_Embed, error) {
	og, err := opengraph.Fetch(event.Issue.GetHTMLURL())
	if err != nil {
		return nil, err
	}

	if len(og.Image) == 0 {
		return nil, fmt.Errorf("no embed data found")
	}

	blobBytes, err := getImage(ctx, og.Image[0].URL)
	if err != nil {
		return nil, err
	}

	blob, err := comatproto.RepoUploadBlob(ctx, b.xrpcc, bytes.NewReader(blobBytes))
	if err != nil {
		return nil, err
	}

	return &appbsky.FeedPost_Embed{
		EmbedExternal: &appbsky.EmbedExternal{
			LexiconTypeID: "",
			External: &appbsky.EmbedExternal_External{
				Description: og.Description,
				Thumb: &util.LexBlob{
					Ref:      blob.Blob.Ref,
					MimeType: "image/png",
				},
				Title: event.Issue.GetTitle(),
				Uri:   og.URL.String(),
			},
		},
	}, nil
}

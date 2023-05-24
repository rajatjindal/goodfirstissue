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
	"github.com/otiai10/opengraph"
)

const (
	base = "https://bsky.social"
)

func NewClient(client *http.Client, handle, password string) (*xrpc.Client, error) {
	xrpcc := &xrpc.Client{
		Auth: &xrpc.AuthInfo{
			Handle: handle,
		},
		Client: client,
		Host:   base,
	}

	auth, err := comatproto.ServerCreateSession(context.TODO(), xrpcc, &comatproto.ServerCreateSession_Input{
		Identifier: xrpcc.Auth.Handle,
		Password:   password,
	})
	if err != nil {
		return nil, err
	}

	xrpcc.Auth.AccessJwt = auth.AccessJwt
	xrpcc.Auth.RefreshJwt = auth.RefreshJwt
	xrpcc.Auth.Did = auth.Did
	xrpcc.Auth.Handle = auth.Handle

	return xrpcc, nil
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

func CreatePost(ctx context.Context, xrpcc *xrpc.Client, msg string, link string) (*comatproto.RepoCreateRecord_Output, error) {
	og, _ := opengraph.Fetch(link)

	postMsg := msg + " \n " + link
	blobBytes, err := getImage(ctx, og.Image[0].URL)
	if err != nil {
		return nil, err
	}

	blob, err := comatproto.RepoUploadBlob(ctx, xrpcc, bytes.NewReader(blobBytes))
	if err != nil {
		return nil, err
	}

	postInp := &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &util.LexiconTypeDecoder{
			&appbsky.FeedPost{
				Text:      postMsg,
				CreatedAt: time.Now().Format(time.RFC3339),
				Facets:    DetectFacets(postMsg),
				Embed: &appbsky.FeedPost_Embed{
					EmbedExternal: &appbsky.EmbedExternal{
						LexiconTypeID: "",
						External: &appbsky.EmbedExternal_External{
							Description: og.Description,
							Thumb: &util.LexBlob{
								Ref:      blob.Blob.Ref,
								MimeType: "blob",
							},
							Title: og.Title,
							Uri:   og.URL.String(),
						},
					},
				},
			},
		},
	}

	return comatproto.RepoCreateRecord(ctx, xrpcc, postInp)
}

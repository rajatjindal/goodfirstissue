package bluesky

import (
	"context"
	"net/http"
	"time"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
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

func CreatePost(ctx context.Context, xrpcc *xrpc.Client, msg string) (*comatproto.RepoCreateRecord_Output, error) {
	postInp := &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &util.LexiconTypeDecoder{
			&appbsky.FeedPost{
				Text:      msg,
				CreatedAt: time.Now().Format(time.RFC3339),
			},
		},
	}

	return comatproto.RepoCreateRecord(ctx, xrpcc, postInp)
}

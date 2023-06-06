package bluesky

import (
	"context"
	"net/http"
	"testing"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/google/go-github/v51/github"
	"github.com/stretchr/testify/require"
)

func TestOne(t *testing.T) {
	defer t.Skip()

	xrpcc := &xrpc.Client{
		Auth: &xrpc.AuthInfo{
			Handle: "goodfirstissue.bsky.social",
		},
		Client: http.DefaultClient,
		Host:   base,
	}

	auth, err := comatproto.ServerCreateSession(context.TODO(), xrpcc, &comatproto.ServerCreateSession_Input{
		Identifier: xrpcc.Auth.Handle,
		Password:   "<enter password>",
	})
	require.Nil(t, err)

	xrpcc.Auth.AccessJwt = auth.AccessJwt
	xrpcc.Auth.RefreshJwt = auth.RefreshJwt
	xrpcc.Auth.Did = auth.Did
	xrpcc.Auth.Handle = auth.Handle

	b := &BlueSky{
		xrpcc: xrpcc,
	}

	err = b.CreatePost(context.TODO(), "some prefix", &github.IssuesEvent{
		Action: ptr("opened"),
		Repo: &github.Repository{
			Name:     ptr("rajatjindal/goodfirstissue"),
			FullName: ptr("rajatjindal/goodfirstissue"),
			Language: ptr("en"),
		},
		Issue: &github.Issue{
			HTMLURL: ptr("https://github.com/rajatjindal/goodfirstissue/issues/44"),
			Labels: []*github.Label{
				{
					Name: ptr("good first issue"),
				},
			},
			Title: ptr("ignore: testing goodfirstissue"),
		},
	})

	require.Nil(t, err)
}

func ptr[T any](v T) *T {
	return &v
}

module github.com/rajatjindal/goodfirstissue

go 1.20

require (
	github.com/dghubble/go-twitter v0.0.0-20221104224141-912508c3888b
	github.com/dghubble/oauth1 v0.7.2
	github.com/fermyon/spin/sdk/go v1.2.1
	github.com/google/go-github/v51 v51.0.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.2
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cloudflare/circl v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dghubble/sling v1.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/sirupsen/logrus v1.9.2 => github.com/rajatjindal/logrus v0.0.0-20230517175950-352781de903c
replace github.com/sirupsen/logrus => ../../sirupsen/logrus

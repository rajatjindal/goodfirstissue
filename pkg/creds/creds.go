package creds

type Provider interface {
	GetCredentials(svc string) (map[string]string, error)
}

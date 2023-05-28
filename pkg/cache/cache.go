package cache

type Provider interface {
	Set(k string, v interface{}) error
	Get(k string) (interface{}, bool)
}

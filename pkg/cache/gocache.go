package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type GoCache struct {
	store *gocache.Cache
}

func NewGoCache(expiration, cleanup time.Duration) *GoCache {
	return &GoCache{
		gocache.New(expiration, cleanup),
	}
}
func (g *GoCache) Get(k string) (interface{}, bool) {
	return g.store.Get(k)
}

func (g *GoCache) Set(k string, v interface{}) error {
	g.store.SetDefault(k, v)
	return nil
}

package cache

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type FakeCache struct {
	mock.Mock
}

var (
	_ Provider = &FakeCache{}
)

func NewFakeCache(expiration, cleanup time.Duration) *FakeCache {
	return &FakeCache{}
}

func (g *FakeCache) Get(k string) (interface{}, bool) {
	args := g.Called(k)
	return args.Get(0), args.Bool(1)
}

func (g *FakeCache) Set(k string, v interface{}) error {
	args := g.Called(k, v)
	return args.Error(0)
}

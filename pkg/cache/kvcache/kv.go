package kvcache

import (
	"encoding/json"
	"time"

	"github.com/fermyon/spin/sdk/go/key_value"
)

const prefix = "cached:"

type kv struct {
	expiration time.Duration
}

func Provider(expiration time.Duration, _ time.Duration) *kv {
	return &kv{}
}

type DataWithExpiry struct {
	Raw       interface{} `json:"raw"`
	ExpiresAt time.Time   `json:"expiresAt"`
}

func (k *kv) Set(key string, value interface{}) error {
	store, err := key_value.Open("default")
	if err != nil {
		return err
	}
	defer key_value.Close(store)

	withExpiry := &DataWithExpiry{
		Raw:       value,
		ExpiresAt: time.Now().Add(k.expiration),
	}

	data, err := json.Marshal(withExpiry)
	if err != nil {
		return err
	}

	return key_value.Set(store, key, data)
}

func (k *kv) Get(key string) (interface{}, bool) {
	store, err := key_value.Open("default")
	if err != nil {
		return nil, false
	}
	defer key_value.Close(store)

	value, err := key_value.Get(store, key)
	if err != nil {
		return nil, false
	}

	withExpiry := DataWithExpiry{}
	err = json.Unmarshal(value, &withExpiry)
	if err != nil {
		return nil, false
	}

	if withExpiry.ExpiresAt.Before(time.Now()) {
		return nil, false
	}

	return withExpiry.Raw, true
}

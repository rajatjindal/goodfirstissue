package cache

import (
	"encoding/json"
	"time"

	"github.com/fermyon/spin/sdk/go/key_value"
)

type DataWithExpiry struct {
	Raw       interface{} `json:"raw"`
	ExpiresAt time.Time   `json:"expiresAt"`
}

func Get(key string) (interface{}, bool) {
	store, err := get_store()
	if err != nil {
		return nil, false
	}

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

func Set(key string, value interface{}) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}

	withExpiry := &DataWithExpiry{
		Raw:       raw,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	data, err := json.Marshal(withExpiry)
	if err != nil {
		return err
	}

	store, err := get_store()
	if err != nil {
		return err
	}

	err = key_value.Set(store, key, data)
	if err != nil {
		return err
	}

	return nil
}

func get_store() (key_value.Store, error) {
	return key_value.Open("default")
}

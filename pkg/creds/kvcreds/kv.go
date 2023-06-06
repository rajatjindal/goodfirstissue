package kvcreds

import (
	"fmt"
	"strings"

	"github.com/fermyon/spin/sdk/go/key_value"
)

type kv struct{}

func Provider() *kv {
	return &kv{}
}

func (k *kv) GetCredentials(svc string) (map[string]string, error) {
	store, err := key_value.Open("default")
	if err != nil {
		return nil, err
	}
	defer key_value.Close(store)

	keys, err := key_value.GetKeys(store)
	if err != nil {
		return nil, err
	}

	credentials := map[string]string{}
	prefix := fmt.Sprintf("%s:", svc)
	for _, key := range keys {
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		value, err := key_value.Get(store, key)
		if err != nil {
			return nil, err
		}

		woutPrefix := strings.TrimPrefix(key, prefix)
		credentials[woutPrefix] = string(value)

	}

	return credentials, nil
}

package kvcache

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fermyon/spin/sdk/go/key_value"
	"github.com/sirupsen/logrus"
)

func (k *kv) CleanupExpiredCache() error {
	store, err := key_value.Open("default")
	if err != nil {
		return err
	}
	defer key_value.Close(store)

	keys, err := key_value.GetKeys(store)
	for _, key := range keys {
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		value, err := key_value.Get(store, key)
		if err != nil {
			logrus.Errorf("error deleting key %s\n", key)
		}

		withExpiry := DataWithExpiry{}
		err = json.Unmarshal(value, &withExpiry)
		if err != nil {
			return err
		}

		if withExpiry.ExpiresAt.After(time.Now()) {
			continue
		}

		err = key_value.Delete(store, key)
		if err != nil {
			logrus.Errorf("error deleting key %s. error: %v\n", key, err)
		}
	}

	return nil
}

func (k *kv) CleanupCache() error {
	store, err := key_value.Open("default")
	if err != nil {
		return err
	}
	defer key_value.Close(store)

	keys, err := key_value.GetKeys(store)
	for _, key := range keys {
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		err = key_value.Delete(store, key)
		if err != nil {
			logrus.Errorf("error deleting key %s\n", key)
		}
	}

	return nil
}

func getKeyWithPrefix(key string) string {
	return fmt.Sprintf("%s%s", prefix, key)
}

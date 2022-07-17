package kvs

import (
	"errors"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	store.RLock()
	result, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return result, nil
}

func Put(key, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

func Del(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}

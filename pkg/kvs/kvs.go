package kvs

import "errors"

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	result, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return result, nil
}

func Put(key, value string) error {
	store[key] = value

	return nil
}

func Del(key string) error {
	delete(store, key)

	return nil
}

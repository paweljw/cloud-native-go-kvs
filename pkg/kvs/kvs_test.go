package kvs

import "testing"
import "github.com/stretchr/testify/assert"

func TestPutGetDel(t *testing.T) {
	err := Put("test", "result")
	assert.Nil(t, err)

	result, ok := store.m["test"]
	assert.Equal(t, "result", result)
	assert.True(t, ok)

	getResult, err := Get("test")
	assert.Equal(t, "result", getResult)
	assert.Nil(t, err)

	err = Del("test")
	assert.Nil(t, err)

	result, ok = store.m["test"]
	assert.Equal(t, "", result)
	assert.False(t, ok)

	getResult, err = Get("test")
	assert.Equal(t, "", getResult)
	assert.Equal(t, ErrorNoSuchKey, err)
}

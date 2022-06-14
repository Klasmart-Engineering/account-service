package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvOrPanic(t *testing.T) {

	f := func(key string) (value string, didPanic bool) {
		defer func() {
			if e := recover(); e != nil {
				didPanic = true
			}
		}()
		value = GetEnvOrPanic(key)
		return value, didPanic
	}

	os.Setenv("test", "123")

	value, didPanic := f("test")
	assert.Equal(t, value, "123")
	assert.Equal(t, didPanic, false)

	value, didPanic = f("random")
	assert.Equal(t, value, "")
	assert.Equal(t, didPanic, true)
}

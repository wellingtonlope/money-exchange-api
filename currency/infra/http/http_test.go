package http

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponseError(t *testing.T) {
	t.Run("should create response error", func(t *testing.T) {
		expectedResponseError := ResponseError{"i'm an error"}
		err := errors.New(expectedResponseError.Message)

		got := NewResponseError(err)

		assert.Equal(t, expectedResponseError, got)
	})

	t.Run("should create an empty response error", func(t *testing.T) {
		got := NewResponseError(nil)

		assert.Empty(t, got)
	})
}

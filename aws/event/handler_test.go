package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestHTTPHeader(t *testing.T) {
	// given
	mesage := "hello world"

	// when
	r, err := handle(mesage)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
	assert.Equal(t, "{\"message\":\"hello world\"}", r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
}
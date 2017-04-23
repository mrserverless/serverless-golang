package main

import (
	"testing"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/stretchr/testify/assert"
)

func TestCustomHeader(t *testing.T) {
	// when
	r, err := handlAPIEvent(&apigatewayproxyevt.Event{}, &runtime.Context{})

	// then
	assert.NoError(t, err)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
}

func TestAPIBody(t *testing.T) {
	// when
	r, err := handlAPIEvent(&apigatewayproxyevt.Event{Body: "input"}, &runtime.Context{})

	// then
	assert.NoError(t, err)
	assert.IsType(t, &Response{}, r)
	//response, _ := r.(*Response)
	assert.Equal(t, &APIBody{"Hello, serverless-golang!", "input"}, r.Body)
}

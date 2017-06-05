package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
)

func TestHandleHTTP(t *testing.T) {
	// given
	todo := "{\"message\":\"hello world\"}"
	evt := &apigatewayproxyevt.Event{Body: todo,}

	// when
	r, err := HandleHTTP(evt, nil)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &apigateway.APIGatewayResponse{}, r)

	apiResponse, _ := r.(*apigateway.APIGatewayResponse)
	assert.Equal(t, http.StatusOK, apiResponse.StatusCode)
	assert.Equal(t, "serverless-golang", apiResponse.Headers["X-Powered-By"])
	assert.Equal(t, "{\"message\":\"hello world\"}", apiResponse.Body)
}

func TestHandleRaw(t *testing.T) {
	// TODO
}

func TestHandle(t *testing.T) {
	// given
	todo := &Todo{"hello world"}

	// when
	r, err := handle(todo)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
	assert.Equal(t, "{\"message\":\"hello world\"}", r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
}
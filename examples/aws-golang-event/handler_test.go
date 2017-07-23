package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
	"github.com/satori/go.uuid"
	"fmt"
	"encoding/json"
	"github.com/yunspace/serverless-golang/examples/aws-golang-event/todo"
)

func TestCreate(t *testing.T) {
	// given
	todo := "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"message\":\"hello world\"}"
	evt := &apigatewayproxyevt.Event{Body: todo,}

	// when
	r, err := Create(evt, nil)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &apigateway.APIGatewayResponse{}, r)

	apiResponse, _ := r.(*apigateway.APIGatewayResponse)
	assert.Equal(t, http.StatusCreated, apiResponse.StatusCode)
	assert.Equal(t, "serverless-golang", apiResponse.Headers["X-Powered-By"])
	assert.Equal(t, todo, apiResponse.Body)
}

func TestGet(t *testing.T) {
	// given
	id := uuid.NewV4()
	evt := &apigatewayproxyevt.Event{
		PathParameters: make(map[string]string),
	}
	evt.PathParameters["id"] = id.String()

	// when
	r, err := Get(evt, nil)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &apigateway.APIGatewayResponse{}, r)

	expected := fmt.Sprintf("{\"id\":\"%s\",\"message\":\"hello world\"}", id.String())

	apiResponse, _ := r.(*apigateway.APIGatewayResponse)
	assert.Equal(t, http.StatusOK, apiResponse.StatusCode)
	assert.Equal(t, "serverless-golang", apiResponse.Headers["X-Powered-By"])
	assert.Equal(t, expected, apiResponse.Body)
}

func TestList(t *testing.T) {
	// given
	evt := &apigatewayproxyevt.Event{}

	// when
	r, err := List(evt, nil)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &apigateway.APIGatewayResponse{}, r)

	apiResponse, _ := r.(*apigateway.APIGatewayResponse)
	assert.Equal(t, http.StatusOK, apiResponse.StatusCode)
	assert.Equal(t, "serverless-golang", apiResponse.Headers["X-Powered-By"])

	var todos []*todo.Todo
	err = json.Unmarshal(apiResponse.Body.([]byte), todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 1)
}
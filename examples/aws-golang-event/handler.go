package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"net/http"
	"encoding/json"
	"log"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
	"github.com/yunspace/serverless-golang/examples/todo"
	"github.com/satori/go.uuid"
)

var todoService todo.TodoService

func init() {
	todoService = todo.NewMockTodoService()
}

func Create(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	t, err := unmarshalEvent(evt)
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}
	err = todoService.Create(t)
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusInternalServerError, err), nil
	}
	return successResponse(http.StatusCreated, t)
}

func Get(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}
	t, err := todoService.Get(id)
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusNotFound, err), nil
	}
	return successResponse(http.StatusOK, t)
}

func List(_ *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {

	todos, _ := todoService.List()

	return successResponse(http.StatusOK, todos)
}

func Update(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}
	t, err := unmarshalEvent(evt)
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}
	t.ID = id
	if err := todoService.Update(t); err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusNotFound, err), nil
	}
	return successResponse(http.StatusOK, t)
}

func Delete(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}

	todoService.Delete(id)
	return successResponse(http.StatusNoContent, nil)
}

func unmarshalEvent(evt *apigatewayproxyevt.Event) (*todo.Todo, error) {
	t := &todo.Todo{}
	err := json.Unmarshal([]byte(evt.Body), t); if err != nil {
		log.Printf("unmarshalling error: %s", err.Error())
		return nil, err
	}
	return t, nil
}

func successResponse(status int, todo interface{}) (*apigateway.APIGatewayResponse, error) {
	response := apigateway.NewAPIGatewayResponseWithBody(status, todo)
	response.Headers["X-Powered-By"] = "serverless-golang"
	return response, nil
}

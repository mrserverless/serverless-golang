package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"net/http"
	"encoding/json"
	"log"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
	"github.com/yunspace/serverless-golang/examples/aws-golang-event/todo"
	"github.com/satori/go.uuid"
)


func Create(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	t, err := unmarshalEvent(evt)
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}
	err = todo.Create(t)
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
	t := todo.Get(id)
	if t == nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusNotFound, nil), nil
	}
	return successResponse(http.StatusOK, t)
}

func List(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {

	todos := todo.List()

	response := apigateway.NewAPIGatewayResponseWithBody(http.StatusOK, todos)
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
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
	todo.Update(t)
	return successResponse(http.StatusOK, t)
}

func Delete(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return apigateway.NewAPIGatewayResponseWithError(http.StatusUnprocessableEntity, err), nil
	}

	todo.Delete(id)
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

func successResponse(status int, todo *todo.Todo) (*apigateway.APIGatewayResponse, error) {
	response := apigateway.NewAPIGatewayResponseWithBody(status, todo)
	response.Headers["X-Powered-By"] = "serverless-golang"
	return response, nil
}

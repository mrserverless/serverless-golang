package apigateway

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"net/http"
	"github.com/yunspace/serverless-golang/api"
	"github.com/satori/go.uuid"
)

type APIGatewayEventUnmarshaler func(event *apigatewayproxyevt.Event) (entity api.Entity, err error)

func HandleCreate(evt *apigatewayproxyevt.Event, unmarshaler APIGatewayEventUnmarshaler, createFunc api.CreateFunc) (*APIGatewayResponse, error) {
	v, err := unmarshaler(evt)
	if err != nil {
		return errResponse(http.StatusUnprocessableEntity, err)
	}
	err = createFunc(v)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err)
	}
	return successResponse(http.StatusCreated, v)
}

func HandleGet(evt *apigatewayproxyevt.Event, getFunc api.GetFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return errResponse(http.StatusUnprocessableEntity, err)
	}
	t, err := getFunc(id)
	if err != nil {
		return errResponse(http.StatusNotFound, err)
	}
	return successResponse(http.StatusOK, t)
}

func HandleList(evt *apigatewayproxyevt.Event, listFunc api.ListFunc) (*APIGatewayResponse, error) {
	todos, err := listFunc(evt.QueryStringParameters)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err)
	}

	return successResponse(http.StatusOK, todos)
}

func HandleUpdate(evt *apigatewayproxyevt.Event, unmarshaler APIGatewayEventUnmarshaler, updateFunc api.UpdateFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return errResponse(http.StatusUnprocessableEntity, err)
	}
	v, err := unmarshaler(evt)
	if err != nil {
		return errResponse(http.StatusUnprocessableEntity, err)
	}
	v.SetID(id)
	if err := updateFunc(v); err != nil {
		return errResponse(http.StatusNotFound, err)
	}
	return successResponse(http.StatusOK, v)
}

func HandleDelete(evt *apigatewayproxyevt.Event, deleteFunc api.DeleteFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return errResponse(http.StatusUnprocessableEntity, err)
	}
	if err := deleteFunc(id); err != nil {
		return errResponse(http.StatusInternalServerError, err)
	}
	return successResponse(http.StatusNoContent, nil)
}

func successResponse(status int, todo interface{}) (*APIGatewayResponse, error) {
	response := NewAPIGatewayResponseWithBody(status, todo)
	response.Headers["X-Powered-By"] = "serverless-golang"
	return response, nil
}

func errResponse(status int, err error) (*APIGatewayResponse, error) {
	return NewAPIGatewayResponseWithError(status, err), nil
}

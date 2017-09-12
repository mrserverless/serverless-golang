package apigateway

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"net/http"
	"github.com/yunspace/serverless-golang/api"
	"github.com/satori/go.uuid"
	"github.com/yunspace/serverless-golang/errors"
)

type APIGatewayEventUnmarshaler func(event *apigatewayproxyevt.Event) (entity api.Entity, err error)

func HandleCreate(evt *apigatewayproxyevt.Event, unmarshaler APIGatewayEventUnmarshaler, createFunc api.CreateFunc) (*APIGatewayResponse, error) {
	v, err := unmarshaler(evt)
	if err != nil {
		return ErrAutoResponse(errors.NewUnprocessableEntityError(err))
	}
	err = createFunc(v)
	if err != nil {
		return ErrAutoResponse(err)
	}
	return SuccessResponse(http.StatusCreated, v)
}

func HandleGet(evt *apigatewayproxyevt.Event, getFunc api.GetFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return ErrAutoResponse(errors.NewUnprocessableEntityError(err))
	}
	t, err := getFunc(id)
	if err != nil {
		return ErrAutoResponse(err)
	}
	return SuccessResponse(http.StatusOK, t)
}

func HandleList(evt *apigatewayproxyevt.Event, listFunc api.ListFunc) (*APIGatewayResponse, error) {
	todos, err := listFunc(evt.QueryStringParameters)
	if err != nil {
		return ErrAutoResponse(err)
	}

	return SuccessResponse(http.StatusOK, todos)
}

func HandleUpdate(evt *apigatewayproxyevt.Event, unmarshaler APIGatewayEventUnmarshaler, updateFunc api.UpdateFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return ErrAutoResponse(errors.NewUnprocessableEntityError(err))
	}
	v, err := unmarshaler(evt)
	if err != nil {
		return ErrAutoResponse(errors.NewUnprocessableEntityError(err))
	}
	v.SetID(id)
	if err := updateFunc(v); err != nil {
		return ErrAutoResponse(err)
	}
	return SuccessResponse(http.StatusOK, v)
}

func HandleDelete(evt *apigatewayproxyevt.Event, deleteFunc api.DeleteFunc) (*APIGatewayResponse, error) {
	id, err := uuid.FromString(evt.PathParameters["id"])
	if err != nil {
		return ErrAutoResponse(errors.NewUnprocessableEntityError(err))
	}
	if err := deleteFunc(id); err != nil {
		return ErrAutoResponse(err)
	}
	return SuccessResponse(http.StatusNoContent, nil)
}

func SuccessResponse(status int, todo interface{}) (*APIGatewayResponse, error) {
	response := NewAPIGatewayResponseWithBody(status, todo)
	response.Headers["X-Powered-By"] = "serverless-golang"
	return response, nil
}

func errResponse(status int, err error) (*APIGatewayResponse, error) {
	return NewAPIGatewayResponseWithError(status, err), nil
}

// ErrAutoResponse follows the logic of asserting for behaviour instead of type
// as per https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
// more behaviour assertions can be added based on needs
func ErrAutoResponse(err error) (*APIGatewayResponse, error) {
	statusCode := http.StatusInternalServerError
	if statusErr, ok := err.(errors.StatusError); ok {
		statusCode = statusErr.Status()
	}

	return errResponse(statusCode, err)
}

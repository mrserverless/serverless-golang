package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"net/http"
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisstreamsevt"
	"log"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
)

type Todo struct {
	Message string `json:"message"`
}

func HandleRaw(evt *json.RawMessage, _ *runtime.Context) (interface{}, error) {
	t := &Todo{}
	err := json.Unmarshal(*evt, t); if err != nil {
		log.Printf("unmarshalling error: %s", err.Error())
		return apigateway.NewAPIGatewayResponse(http.StatusUnprocessableEntity), nil
	}

	resp := apigateway.NewAPIGatewayResponse(http.StatusOK)
	resp.SetBody(evt)
	return resp, nil
}

func HandleHTTP(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	t := &Todo{}
	err := json.Unmarshal([]byte(evt.Body), t); if err != nil {
		log.Printf("unmarshalling error: %s", err.Error())
		return apigateway.NewAPIGatewayResponse(http.StatusUnprocessableEntity), nil
	}
	return handle(t)
}

func HandleKinesis(evt kinesisstreamsevt.Event, _ *runtime.Context) (interface{}, error) {
	for _, record := range evt.Records {
		log.Printf("received: %s", record.String())
	}
	return nil, nil
}

func handle(todo *Todo) (*apigateway.APIGatewayResponse, error) {
	response := apigateway.NewAPIGatewayResponse(http.StatusOK)
	response.SetBody(todo)
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
}

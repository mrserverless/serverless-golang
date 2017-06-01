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
	bytes, err := evt.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return handle(string(bytes))
}

func HandleHTTP(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return handle(evt.Body)
}

func HandleKinesis(evt kinesisstreamsevt.Event, _ *runtime.Context) (interface{}, error) {
	for _, record := range evt.Records {
		log.Printf("received: %s", record.String())
	}
	return nil, nil
}

func handle(body string) (*apigateway.APIGatewayResponse, error) {
	response := &apigateway.APIGatewayResponse{
		StatusCode: http.StatusOK,
		Headers:    make(map[string]string),
	}
	response.SetBody(body)
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
}

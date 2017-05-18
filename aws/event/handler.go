package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"net/http"
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisstreamsevt"
	"log"
)

func HandleRaw(evt *json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	bytes, err := evt.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return handle(string(bytes))
}

func HandleHTTP(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	return handle(evt.Body)
}

func HandleKinesis(evt kinesisstreamsevt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, record := range evt.Records {
		log.Printf("received: %s", record.String())
	}
	return nil, nil
}

func handle(body string) (*HTTPResponse, error) {
	response := &HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    make(map[string]string),
	}
	response.SetBody(&Todo{body})
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
}

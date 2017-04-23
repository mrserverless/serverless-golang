package main

import (
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"log"
)

func handlAPIEvent(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (*Response, error) {
	log.Println(evt)
	response := &Response{
		StatusCode: 200,
		Headers:    make(map[string]string),
	}
	response.SetBody(&APIBody{"Hello, serverless-golang!", evt.Body})
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
}

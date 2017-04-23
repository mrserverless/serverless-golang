package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func EventHandle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	return handlAPIEvent(evt, ctx)
}

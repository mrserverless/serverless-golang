package main

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/yunspace/serverless-golang/aws/net"
	"github.com/yunspace/serverless-golang/aws/event"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Handle is the exported handler called by AWS Lambda.
var NetHandle apigatewayproxy.Handler

func init() {
	NetHandle = net.NewNetHandler()
}

func EventHandle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	return event.HandlHTTPEvent(evt, ctx)
}
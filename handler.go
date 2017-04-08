package main

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/yunspace/serverless-golang/net"
	"github.com/yunspace/serverless-golang/shim"
)

// Handle is the exported handler called by AWS Lambda.
var NetHandle apigatewayproxy.Handler

func init() {
	NetHandle = net.NewHTTPHandler()
}

func EventHandle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return shim.HandleEvent(evt, ctx)
}
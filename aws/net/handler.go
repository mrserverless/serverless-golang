package main

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
)

// Handle is the exported handler called by AWS Lambda.
var NetHandle apigatewayproxy.Handler

func init() {
	NetHandle = NewNetHandler()
}

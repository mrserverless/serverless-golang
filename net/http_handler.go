package net

import (
	"net/http"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
)

func NewHTTPHandler() apigatewayproxy.Handler {
	ln := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	handle := apigatewayproxy.New(ln, nil).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
	go http.Serve(ln, http.HandlerFunc(handleHTTP))

	return handle
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Powered-By", "serverless-golang")
	w.Write([]byte("Hello, serverless-golang!"))
}
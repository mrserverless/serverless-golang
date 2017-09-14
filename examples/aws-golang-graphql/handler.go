package main

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"

	"net/http"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

// Handler is the exported handler called by AWS Lambda.
var Handler apigatewayproxy.Handler
var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(Schema, &Resolver{})

	Handler = NewHandler()
}

func NewHandler() apigatewayproxy.Handler {
	ln := net.Listen()

	handle := apigatewayproxy.New(ln, nil).Handle

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	go http.Serve(ln, nil)

	return handle
}

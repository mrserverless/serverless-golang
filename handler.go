package main

import "C"

import (
	"encoding/json"
	"net/http"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return NewResponse(http.StatusOK, NewBody("Hello Serverless Golang!", "empty input")), nil
}

package main

import "C"

import (
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return "Hello, World!", nil
}

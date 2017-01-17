package main

import "C"

import (
	"encoding/json"
	"net/http"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	r := Response{StatusCode: http.StatusOK}
	r.SetBodyStruct(&Body{Message: "Hello Serverless Golang!", Input: evt})
	return r, nil
}

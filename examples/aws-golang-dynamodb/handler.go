package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/yunspace/serverless-golang/aws/event/apigateway"
	"github.com/yunspace/serverless-golang/examples/todo"
	"github.com/yunspace/serverless-golang/api"
	"github.com/yunspace/serverless-golang/examples/aws-golang-dynamodb/services"
	"github.com/yunspace/serverless-golang/examples/aws-golang-dynamodb/config"
)

var todoDynamoService api.CRUDAPI

func init() {
	todoDynamoService = services.NewTodoDynamoDBService(config.NewConfig().DB)
}

/// Create
func Create(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return apigateway.HandleCreate(evt, todo.TodoAPIGAtewayEventUnmarshaler, todoDynamoService.Create)
}

func Get(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return apigateway.HandleGet(evt, todoDynamoService.Get)
}

func List(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return apigateway.HandleList(evt, todoDynamoService.List)
}

func Update(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return apigateway.HandleUpdate(evt, todo.TodoAPIGAtewayEventUnmarshaler, todoDynamoService.Update)
}

func Delete(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	return apigateway.HandleDelete(evt, todoDynamoService.Delete)
}
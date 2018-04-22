package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/yunspace/serverless-golang/examples/aws-dynamo-todo/config"
	"github.com/yunspace/serverless-golang/examples/aws-dynamo-todo/services"
	"github.com/yunspace/serverless-golang/examples/todo"
)

var todoDynamoService api.CRUDAPI

func init() {
	todoDynamoService = services.NewTodoDynamoDBService(config.NewConfig().DB)
}

func Create(evt *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apigateway.HandleCreate(evt, todo.TodoAPIGAtewayEventUnmarshaler, todoDynamoService.Create)
}

func Get(evt *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apigateway.HandleGet(evt, todoDynamoService.Get)
}

func List(evt *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apigateway.HandleList(evt, todoDynamoService.List)
}

func Update(evt *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apigateway.HandleUpdate(evt, todo.TodoAPIGAtewayEventUnmarshaler, todoDynamoService.Update)
}

func Delete(evt *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apigateway.HandleDelete(evt, todoDynamoService.Delete)
}

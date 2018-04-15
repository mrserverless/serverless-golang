package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/yunspace/serverless-golang/examples/aws-kinesis/config"
	"github.com/yunspace/serverless-golang/examples/aws-kinesis/services"
)

var kinesisServices *services.KinesisService

func DataPublisher(evt json.RawMessage) (interface{}, error) {
	bytes, err := evt.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return kinesisServices.Publish(bytes)
}

func main() {
	kinesisServices = services.NewKinesisService(config.NewConfig())
	lambda.Start(Handler)
}

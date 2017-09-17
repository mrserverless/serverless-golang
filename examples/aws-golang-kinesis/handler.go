package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisstreamsevt"
	"log"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/yunspace/serverless-golang/examples/aws-golang-kinesis/services"
	"github.com/yunspace/serverless-golang/examples/aws-golang-kinesis/config"
	"encoding/json"
)

var kinesisServices *services.KinesisService

func init() {
	kinesisServices = services.NewKinesisService(config.NewConfig())
}

func DataPublisher(evt json.RawMessage, _ *runtime.Context) (interface{}, error) {
	bytes, err := evt.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return kinesisServices.Publish(bytes)
}

func DataLogger (evt *kinesisstreamsevt.Event, _ *runtime.Context) (interface{}, error) {
	rs := evt.Records
	for i, r := range rs {
		log.Printf("Record %d: %s\n", i, r)
	}

	return evt.String(), nil
}
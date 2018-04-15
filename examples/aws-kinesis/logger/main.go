package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func DataLogger(evt *events.KinesisEvent) (interface{}, error) {
	rs := evt.Records
	for i, r := range rs {
		log.Printf("Record %d: %s\n", i, r)
	}

	return evt.String(), nil
}

func main() {
	lambda.Start(Handler)
}

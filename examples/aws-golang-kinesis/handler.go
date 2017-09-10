package aws_golang_kinesis

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisstreamsevt"
	"log"
)

func DataLogger (evt *kinesisstreamsevt.Event) (interface{}, error) {
	rs := evt.Records
	for r := range rs {
		log.Println(r)
	}

	return evt.String(), nil
}
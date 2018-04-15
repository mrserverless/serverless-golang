package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type ProtobufHandler func(context.Context, []byte) ([]byte, error)

func (handler ProtobufHandler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	// unmarshal Protobuf

	return handler(ctx, payload)
}

var _ lambda.Handler = (ProtobufHandler)(nil)

func NewProtobufHandler() ProtobufHandler {
	return func(ctx context.Context, request []byte) ([]byte, error) {
		return nil, nil
	}
}

func main() {
	lambda.StartHandler(NewProtobufHandler())
}

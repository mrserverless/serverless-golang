package shim

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func HandleEvent(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return nil, nil
}


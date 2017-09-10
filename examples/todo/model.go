package todo

import (
	"github.com/satori/go.uuid"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/yunspace/serverless-golang/api"
	"encoding/json"
)

type Todo struct {
	Id      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

func (t *Todo) ID() [16]byte {
	return t.Id
}

func (t *Todo) SetID(id [16]byte) {
	t.Id = id
}

func TodoAPIGAtewayEventUnmarshaler(evt *apigatewayproxyevt.Event) (api.Entity, error) {
	t := &Todo{}
	if err := json.Unmarshal([]byte(evt.Body), t); err != nil {
		return nil, err
	}
	return t, nil
}


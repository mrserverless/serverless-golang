package apigateway

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/satori/go.uuid"
	"fmt"
	"encoding/json"
	"github.com/yunspace/serverless-golang/examples/todo"
	"github.com/yunspace/serverless-golang/api"
)

func givenCRUD() api.CRUDAPI {
	return todo.NewMockTodoService()
}

func TestCreate_return200AndBodyWhenSuccess(t *testing.T) {
	// given
	id := uuid.NewV4()
	givenTodo := fmt.Sprintf(`{
		"id": "%s",
		"message":"hello world"
	}`, id)
	evt := &apigatewayproxyevt.Event{Body: givenTodo}

	// when
	r, err := HandleCreate(evt, todo.TodoAPIGAtewayEventUnmarshaler, givenCRUD().Create)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &APIGatewayResponse{}, r)

	assert.Equal(t, http.StatusCreated, r.StatusCode)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
	assert.JSONEq(t, givenTodo, r.Body.(string))
}

func TestGet_return200AndBodyWhenSuccess(t *testing.T) {
	// given
	id := uuid.NewV4()
	evt := &apigatewayproxyevt.Event{
		PathParameters: make(map[string]string),
	}
	evt.PathParameters["id"] = id.String()

	// when
	r, err := HandleGet(evt, givenCRUD().Get)

	// then
	assert.NoError(t, err)
	assert.IsType(t, &APIGatewayResponse{}, r)

	expected := fmt.Sprintf("{\"id\":\"%s\",\"message\":\"hello world\"}", id.String())

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
	assert.JSONEq(t, expected, r.Body.(string))
}

func TestList_return200AndBodiesWhenSuccess(t *testing.T) {
	// given
	evt := &apigatewayproxyevt.Event{}
	crudService := givenCRUD()
	crudService.Create(&todo.Todo{Id: uuid.NewV4()})
	crudService.Create(&todo.Todo{Id: uuid.NewV4()})
	crudService.Create(&todo.Todo{Id: uuid.NewV4()})

	// when
	r, err := HandleList(evt, crudService.List)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])

	var todos []*todo.Todo
	err = json.Unmarshal([]byte(r.Body.(string)), &todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 3)
}

func TestUpdate_return200AndBodyWhenSuccess(t *testing.T) {
	// given
	id := uuid.NewV4()

	crudService := givenCRUD()
	crudService.Create(&todo.Todo{Id: id, Message: "I need an Update!"})
	updatedTodo := fmt.Sprintf(`{
		"id": "%s",
		"message":"I got an Update!'"
	}`, id)

	// when
	evt := &apigatewayproxyevt.Event{
		PathParameters: make(map[string]string),
		Body:           updatedTodo,
	}
	evt.PathParameters["id"] = id.String()
	r, err := HandleUpdate(evt, todo.TodoAPIGAtewayEventUnmarshaler, crudService.Update)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])
	assert.JSONEq(t, updatedTodo, r.Body.(string))
}

func TestUpdate_return204WhenSuccess(t *testing.T) {
	// given
	id := uuid.NewV4()

	crudService := givenCRUD()
	crudService.Create(&todo.Todo{Id: id, Message: "I will be deleted"})

	// when
	evt := &apigatewayproxyevt.Event{
		PathParameters: make(map[string]string),
	}
	evt.PathParameters["id"] = id.String()
	r, err := HandleDelete(evt, crudService.Delete)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, r.StatusCode)
	assert.Equal(t, "serverless-golang", r.Headers["X-Powered-By"])

	todos, _ := crudService.List(nil)
	assert.Len(t, todos, 0)
}

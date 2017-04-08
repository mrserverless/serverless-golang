package event

import (
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

func HandlHTTPEvent(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	response := &Response {
		StatusCode: 200,
		Body: "Hello, serverless-golang!",
	}
	response.Headers["X-Powered-By"] = "serverless-golang"

	return response, nil
}


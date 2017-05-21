package apigateway

import "encoding/json"

type APIGatewayResponse struct {
	StatusCode    int                 `json:"statusCode"`
	Headers       map[string]string   `json:"headers"`
	Body          interface{}         `json:"body"`
	Base64Encoded bool                `json:"isBase64Encoded"`
}

// inspired by serverless-java
func (r *APIGatewayResponse) SetBody(b interface{}) {
	bytes, _ := json.Marshal(b)
	r.Body = string(bytes)
}

package apigateway

import "encoding/json"

type APIGatewayResponse struct {
	StatusCode    int               `json:"statusCode"`
	Headers       map[string]string `json:"headers"`
	Body          interface{}       `json:"body"`
	Base64Encoded bool              `json:"isBase64Encoded"`
}

func NewAPIGatewayResponse(status int) *APIGatewayResponse {
	return &APIGatewayResponse{
		StatusCode:    status,
		Base64Encoded: false,
		Headers:       make(map[string]string),
	}
}

func NewAPIGatewayResponseWithBody(status int, body interface{}) *APIGatewayResponse {
	r := NewAPIGatewayResponse(status)
	r.SetBody(body)
	return r
}

func NewAPIGatewayResponseWithError(status int, error error) *APIGatewayResponse {
	return &APIGatewayResponse{
		StatusCode:    status,
		Base64Encoded: false,
		Headers:       make(map[string]string),
		Body:          error.Error(),
	}
}

// SetBody checks to see if body implements json.Marshaler first
// If it does, it will use MarshalJSON() function.
// If not, defer back to json.MarshalIndent()
func (r *APIGatewayResponse) SetBody(b interface{}) {
	if m, ok := b.(json.Marshaler); ok {
		bytes, _ := m.MarshalJSON()
		r.Body = string(bytes)
		return
	}

	bytes, _ := json.MarshalIndent(b, "", "")
	r.Body = string(bytes)
}
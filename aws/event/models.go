package main

import "encoding/json"


// special format need by API Gateway response body
type Todo struct {
	Message string `json:"message"`
}

type HTTPResponse struct {
	StatusCode int          	`json:"statusCode"`
	Headers    map[string]string  	`json:"headers"`
	Body	   interface{}          `json:"body"`
}

// inspired by serverless-java
func (r *HTTPResponse) SetBody(b interface{}) {
	bytes, _ := json.Marshal(b)
	r.Body = string(bytes)
}
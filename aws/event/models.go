package main

import "encoding/json"

// special format need by API Gateway response body
type APIBody struct {
	Message string `json:"message"`
	Input   string `json:"input"`
}

type Response struct {
	StatusCode int          	`json:"statusCode"`
	Headers    map[string]string  	`json:"headers"`
	Body	   string               `json:"body"`
}

// inspired by serverless-java
func (r *Response) SetBody(b *APIBody) {
	bytes, _ := json.Marshal(b)
	r.Body = string(bytes)
}
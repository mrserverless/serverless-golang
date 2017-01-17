package main

import "encoding/json"

type Body struct {
	Message string		`json:"message"`
	Input   json.RawMessage	`json:"input"`
}

type Response struct {
	StatusCode int		`json:"statusCode"`
	Body       string	`json:"body"`
}

// inspired by serverless-java
func (r *Response) SetBodyStruct(b *Body) {
	bytes, _ := json.Marshal(b)
	r.Body = string(bytes)
}
package main

type Body struct {
	Message string `json:"message"`
	Input   string `json:"input"`
}

func NewBody(m string, e string) *Body {
	return &Body{Message: m, Input: e}
}

type Response struct {
	StatusCode int   `json:"statusCode"`
	Body       *Body `json:"body"`
}

func NewResponse(s int, b *Body) *Response {
	return &Response{StatusCode: s, Body: b}
}

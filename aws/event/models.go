package main

// special format need by API Gateway response body
type APIBody struct {
	Message string `json:"message"`
	Input   string `json:"input"`
}

type Response struct {
	StatusCode int          	`json:"statusCode"`
	Headers    map[string]string  	`json:"headers"`
	Body	   *APIBody                `json:"body"`
}
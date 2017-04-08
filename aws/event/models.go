package event

type Body struct {
	Message string `json:"message"`
	Input   string `json:"input"`
}

func NewBody(m string, e string) *Body {
	return &Body{Message: m, Input: e}
}

type Response struct {
	StatusCode int          	`json:"statusCode"`
	Headers    map[string]string  	`json:"headers"`
	Body       string 		`json:"body"`
}
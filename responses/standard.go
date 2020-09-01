package responses

// Standard describes standard rendered JSON for Standard Response
type Standard struct {
	RequestParam string      `json:"request_param"`
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
	Next         string      `json:"next"`
	Version      Version     `json:"version"`
}

// Version : describes standard rendered JSON for Apps Versioning
type Version struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

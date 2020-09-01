package logger

//LogFormat ....
type LogFormat struct {
	HTTPMethod string      `json:"method"`
	URL        string      `json:"url"`
	Host       string      `json:"host"`
	Status     int         `json:"status"`
	Response   interface{} `json:"response"`
	Request    interface{} `json:"request"`
	Message    string      `json:"message"` //message for stackdriver
}

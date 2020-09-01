package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aeeem/LogStack/logger"
)

type (

	//customResponseWriter ...
	customResponseWriter struct {
		http.ResponseWriter
		Request    json.RawMessage
		Response   interface{}
		StatusCode int
	}
)

type Default struct{}

//DefaultLog ...
func (d *Default) DefaultLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoveryFromPanic(w, r)
		lw := &customResponseWriter{
			ResponseWriter: w,
		}

		log := ExtractRequest(r)
		h.ServeHTTP(lw, r)
		log.Status = lw.StatusCode
		log.Response = lw.Response
		log.Log("info", os.Getenv("APP_NAME"))
	})
}

//WriteHeader is a function that overiding default WriteHeader
func (w *customResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *customResponseWriter) Write(data []byte) (int, error) {
	_ = json.Unmarshal(data, &w.Response)
	return w.ResponseWriter.Write(data)
}

//ExtractRequest ..
func ExtractRequest(r *http.Request) logger.LogFormat {
	buff, err := ioutil.ReadAll(r.Body)
	request1 := ioutil.NopCloser(bytes.NewBuffer(buff))
	requestBody := ioutil.NopCloser(bytes.NewBuffer(buff))
	r.Body = request1
	if err != nil {
		var logs logger.NewLogger
		logs.WithFields(map[string]interface{}{"error": err})
		return logger.LogFormat{}
	}
	decoder := json.NewDecoder(requestBody)

	var temp map[string]interface{}
	_ = decoder.Decode(&temp)
	host, _ := os.Hostname()
	url := r.URL.RequestURI()
	return logger.LogFormat{
		Host:       host,
		HTTPMethod: r.Method,
		URL:        url,
		Message:    "API Log for url : " + url,
		Request: map[string]interface{}{
			"query":  r.URL.Query,
			"body":   temp,
			"header": r.Header,
		},
	}
}

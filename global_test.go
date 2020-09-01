package global_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	global "github.com/aeeem/LogStack"
	"github.com/gorilla/mux"
)

func TestLogsWarn(t *testing.T) {
	global.InitLogs("none", "none", "dekape-dev-server")
	test := global.Logs

	test.Warn("just test")
	test.WithFields(map[string]interface{}{
		"test_hehe": "ini test",
	}).Warn("test")
}

func TestLogsInfo(t *testing.T) {
	global.InitLogs("none", "none", "dekape-dev-server")
	test := global.Logs
	test.Info("just test")
	test.WithFields(map[string]interface{}{
		"test_hehe": "ini test",
	}).Info("test")
}

func TestLogsError(t *testing.T) {
	global.InitLogs("none", "none", "dekape-dev-server")
	test := global.Logs

	test.Error("just test")
	test.WithFields(map[string]interface{}{
		"test_hehe": "ini test",
	}).Error("test")
}

func TestLogsDebug(t *testing.T) {
	global.InitLogs("JSON", "none", "dekape-dev-server")
	test := global.Logs

	test.Debug("just test")
	test.WithFields(map[string]interface{}{
		"test_hehe": "ini test",
	}).Debug("test")
}

func TestLogsDebugWithStruct(t *testing.T) {
	global.InitLogs("json", "none", "dekape-dev-server")
	test := global.Logs

	test.DebugWithStruct("just test", struct {
		Name   string
		Number int
	}{
		Name:   "Nama",
		Number: 787878,
	})
	test.WithFields(map[string]interface{}{
		"test_hehe": "ini test",
	}).Debug("test")
}

func TestLogMiddleware(t *testing.T) {
	global.InitLogs("json", "none", "dekape-dev-server")
	router := mux.NewRouter()
	router.Use(global.Middleware.DefaultLog)
	router.HandleFunc("/test", DummyController).Methods("POST")
	request, _ := json.Marshal(map[string]interface{}{
		"test": "body",
	})
	req, _ := http.NewRequest("POST", "http://localhost:8000/test", bytes.NewBuffer(request))
	req.Header.Add("HAHA", "hehe")
	router.ServeHTTP(httptest.NewRecorder(), req)

}

func DummyController(w http.ResponseWriter, r *http.Request) {
	body := map[string]interface{}{}
	var header interface{}
	json.NewDecoder(r.Body).Decode(&body)
	header = r.Header
	log.Print(header)
	log.Print(body)

}

func DummyRouter() {

}

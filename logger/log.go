package logger

import (
	"encoding/json"

	"cloud.google.com/go/logging"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//WithFields Its consuming Map string interface{}
func (e *NewLogger) WithFields(Fields map[string]interface{}) *NewLogger {
	e.data = Fields
	return e
}

//NewLogger is initiate  new logger in this package
type NewLogger struct {
	data map[string]interface{}
}

//JSONLog set formatter into json
func JSONLog() {
	log.Formatter = new(logrus.JSONFormatter)
}

func toMapString(d interface{}) logrus.Fields {
	b, _ := json.Marshal(d)
	var data logrus.Fields
	_ = json.Unmarshal(b, &data)
	return data
}

//Log is default logging for http  ..
func (e *LogFormat) Log(Type string, Message string) {
	data := log.WithFields(toMapString(e))
	var Severity logging.Severity
	switch Type {
	case "error":
		log.Level = logrus.ErrorLevel
		Severity = logging.Error
		data.Error(Message)
	case "info":
		log.Level = logrus.InfoLevel
		Severity = logging.Info
		data.Info(Message)

	case "warning":
		log.Level = logrus.WarnLevel
		Severity = logging.Warning
		data.Warn(Message)

	default:
		log.Level = logrus.DebugLevel
		Severity = logging.Debug
		data.Debug(Message)
	}
	StackData(data, Severity)

}

//Info is standard log for info
func (e *NewLogger) Info(message string) {
	if e.data != nil {
		log.Level = logrus.InfoLevel
		fields := logrus.Fields{}
		fields = e.data
		entry := log.WithFields(fields)
		entry.Info(message)
		StackData(entry, logging.Info)
	}
	StackDataBasic(message, logging.Debug)
}

//Debug is standard log for Debug
func (e *NewLogger) Debug(message string) {
	log.Level = logrus.DebugLevel
	maps := make(map[string]interface{})
	maps["message"] = message
	maps["data"] = e.data
	fields := logrus.Fields{}
	e.data = maps
	fields = e.data
	entry := log.WithFields(fields)
	entry.Debug(message)
	StackData(entry, logging.Debug)
}

//Warn is standard log for Warn
func (e *NewLogger) Warn(message string) {
	if e.data != nil {
		log.Level = logrus.WarnLevel
		fields := logrus.Fields{}
		fields = e.data
		entry := log.WithFields(fields)
		entry.Warn(message)
		StackData(entry, logging.Warning)
	}
	StackDataBasic(message, logging.Warning)
}

//Error is standard log for Error
func (e *NewLogger) Error(message string) {
	if e.data != nil {
		log.Level = logrus.ErrorLevel
		fields := logrus.Fields{}
		fields = e.data
		entry := log.WithFields(fields)
		entry.Error(message)
		StackData(entry, logging.Error)
	}
	StackDataBasic(message, logging.Error)

}

//OtherTest is standard log for Error
func (e *NewLogger) OtherTest(message string) {
	log.Level = logrus.DebugLevel
	fields := logrus.Fields{}
	fields = e.data
	entry := log.WithFields(fields)
	entry.Error(message)
	StackData(entry, logging.Debug)
}

//ErrorWithStruct logs message with struct (shorthand for withfields)
func (e *NewLogger) ErrorWithStruct(message string, structure interface{}) {
	e.WithFields(
		map[string]interface{}{
			"struct_log": structure,
		},
	).Error(message)

	//	log.Level = logrus.ErrorLevel
	//	fields := logrus.Fields{message: structure}
	//	entry := log.WithFields(fields)
	//	entry.Error(message)
	//	StackData(entry, logging.Error)
}

//DebugWithStruct logs message with struct (shorthand for withfields)
func (e *NewLogger) DebugWithStruct(message string, structure interface{}) {
	e.WithFields(
		map[string]interface{}{
			"struct_log": structure,
		},
	).Debug(message)
	//	log.Level = logrus.DebugLevel
	//	fields := logrus.Fields{message: structure}
	//	entry := log.WithFields(fields)
	//	entry.Debug(message)
	//	StackData(entry, logging.Debug)
}

//Print is standard log for Error
func (e *NewLogger) Print(data ...interface{}) {
	log.Level = logrus.DebugLevel
	fields := logrus.Fields{"message": "GORM SQL Query Log", "module": "gorm", "data": data}
	entry := log.WithFields(fields)
	entry.Print()
	StackData(entry, logging.Debug)
}

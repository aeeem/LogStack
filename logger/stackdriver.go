package logger

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"cloud.google.com/go/logging"
	"github.com/sirupsen/logrus"
)

//Logs for init global var logger
var Logs NewLogger
var hooksType string
var projectID string

//SetHooks for setting hooks based on main threw
func SetHooks(typ string, projectid string) {
	hooksType = typ
	projectID = projectid
	log.Print(hooksType)
}

func StackDataBasic(entry string, Severity logging.Severity) {
	hooks := strings.Contains(hooksType, "STACK")
	if hooks {
		ctx := context.Background()
		client, err := logging.NewClient(ctx, projectID)
		if err != nil {
			logrus.Warning("Failed to create driverClient. Are you online? ", err)
			return
		}
		defer client.Close()
		//Sanitize input because stackdriver only support _ - alphabet and numeric and no spaces
		reg, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
		logName := reg.ReplaceAllString(os.Getenv("APP_NAME"), "")
		lg := client.Logger(logName)

		//stackdriver logging entry has already automagically determine the type of payload, if string, it will be inputted to textPayload, if struct it will be inputted to jsonPayload and marshalled internally
		entrys := logging.Entry{
			Payload:  entry,
			Severity: Severity,
		}
		lg.Log(entrys)
		return
	}
}

//StackData logging to stack driver
func StackData(entry *logrus.Entry, Severity logging.Severity) {
	hooks := strings.Contains(hooksType, "STACK")
	if hooks {
		ctx := context.Background()
		client, err := logging.NewClient(ctx, projectID)
		if err != nil {
			logrus.Warning("Failed to create driverClient. Are you online? ", err)
			return
		}
		defer client.Close()
		//Sanitize input because stackdriver only support _ - alphabet and numeric and no spaces
		reg, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
		logName := reg.ReplaceAllString(os.Getenv("APP_NAME"), "")
		lg := client.Logger(logName)
		b, _ := json.Marshal(entry.Data)
		entrys := logging.Entry{
			Payload:  json.RawMessage(b),
			Severity: Severity,
		}
		lg.Log(entrys)
		return
	}

}

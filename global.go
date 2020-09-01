package global

import (
	"log"
	"os"
	"strings"

	"github.com/aeeem/LogStack/logger"
	"github.com/aeeem/LogStack/middleware"
)

//Logs is like log.print() with different level of logging
var Logs logger.NewLogger

//Middleware is for default Logging Middleware
var Middleware middleware.Default

//InitLogs is initialize logging for a logger
func InitLogs(format string, hookstype string, projectid string) {
	if strings.Contains(format, "JSON") {
		Logs.Info("Setting up logs into JSON")
		logger.JSONLog()
		Logs.Info("Done")
	}
	log.Print(os.Getenv("LOG_HOOKS"))
	logger.SetHooks(os.Getenv("LOG_HOOKS"), os.Getenv("PROJECT_ID"))
}

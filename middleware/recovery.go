package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"gitlab.com/dekape/dekalogs/logger"
)

//RecoveryFromPanic ...
func RecoveryFromPanic(w http.ResponseWriter, r *http.Request) {
	if rec := recover(); rec != nil {
		var file []string
		i := 0
		for {
			_, tempFile, tempLine, status := runtime.Caller(i)
			file = append(file, tempFile+" : "+fmt.Sprint(tempLine))
			if !status {
				break
			}
			i++
		}
		log := ExtractRequest(r)
		log.Response = map[string]interface{}{
			"errors": rec,
			"trace":  file,
		}
		log.Status = 500
		log.Log("error", "Unexpected!")
		logger.ResponseRender("error", w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

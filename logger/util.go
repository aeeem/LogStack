package logger

import (
	"fmt"
	"net/http"

	"gitlab.com/dekape/dekalogs/responses"

	"github.com/thedevsaddam/renderer"
)

// Rnd ...
var Rnd = renderer.New()

// ResponseRender ...
func ResponseRender(types string, w http.ResponseWriter, statusCode int, v interface{}) {
	if types == "error" {
		responseData := responses.Standard{
			RequestParam: "",
			Status:       "error",
			ErrorMessage: fmt.Sprintf("%v", v),
			Data:         nil,
			Next:         "",
			Version:      responses.Version{Code: "1", Name: "0.1.0"},
		}
		Rnd.JSON(w, statusCode, responseData)
	} else {
		responseData := responses.Standard{
			RequestParam: "",
			Status:       "success",
			ErrorMessage: "",
			Data:         v,
			Next:         "",
			Version:      responses.Version{Code: "1", Name: "0.1.0"},
		}
		Rnd.JSON(w, statusCode, responseData)
	}
}

// // TimeLocalNow ...
// func TimeLocalNow() *time.Time {
// 	loc, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		Info(err)
// 	}
// 	now := time.Now().In(loc)
// 	return &now
// }

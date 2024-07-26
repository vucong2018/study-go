package respone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponeData struct {
	Code 	int 		`json:"code"`
	Message string 		`json:"message"`
	Data 	interface{} `json:"data"` 
}




//success respone

func SuccessRespone(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponeData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

// error

func ErrorRespone(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponeData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}
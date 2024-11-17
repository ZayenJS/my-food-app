package appHttp

import (
	"fmt"
	"runtime"

	"github.com/ZayenJS/config"
	"github.com/gin-gonic/gin"
)

type Response struct {
	context *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{context: c}
}

func (response *Response) JSON(status int, data interface{}) {
	response.context.JSON(status, data)
}

func (response *Response) Error(status int, err error) {
	isDevMode := config.Get("ENV") == "development"

	data := make(map[string]interface{})
	errorMessage := "An unexpected error occurred"

	switch status {
	case 404:
		errorMessage = "Resource not found"
	case 500:
		errorMessage = "Internal server error"
	}

	data["error"] = errorMessage

	if isDevMode {
		data["error"] = err.Error()
		data["trace"] = printStackTrace()
	}

	response.context.JSON(status, data)
}

func printStackTrace() string {
	// Create a byte slice to hold the stack trace
	stackTrace := make([]byte, 1024)
	// Capture the stack trace
	n := runtime.Stack(stackTrace, true)
	// Print the stack trace
	return fmt.Sprintf("Stack trace:%s", string(stackTrace[:n]))
}
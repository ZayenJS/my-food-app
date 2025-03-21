package appHttp

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/ZayenJS/appFs"
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
	env := config.Get("ENV")
	isDevMode := env == "development" || env == "dev"

	data := make(map[string]interface{})
	errorMessage := "An unexpected error occurred"

	switch status {
	case 400:
		errorMessage = "Bad request"
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

func (response *Response) SaveUploadedFile(fieldName string, path string, fileName string) (string, error) {
	file, err := response.context.FormFile(fieldName)

	if err != nil {
		return "", err
	}
	ext := strings.Split(file.Filename, ".")[1]
	fullPath := fmt.Sprintf("%s/%s%s", path, fileName, ext)
	appFs.EnsureDirectoryExists(path)

	err = response.context.SaveUploadedFile(file, fullPath)

	return fullPath, err
}

func printStackTrace() string {
	// Create a byte slice to hold the stack trace
	stackTrace := make([]byte, 1024)
	// Capture the stack trace
	n := runtime.Stack(stackTrace, true)
	// Print the stack trace
	return fmt.Sprintf("Stack trace:%s", string(stackTrace[:n]))
}

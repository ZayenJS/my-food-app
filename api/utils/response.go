package response

import (
	"github.com/ZayenJS/config"
	"github.com/gin-gonic/gin"
)

type Response struct {
	context *gin.Context
}

func New(c *gin.Context) *Response {
	return &Response{context: c}
}

func (response *Response) JSON(status int, data interface{}) {
	response.context.JSON(status, data)
}

func (response *Response) Error(status int, err error) {
	isDevMode := config.Get("ENV") == "development"

	switch status {
	case 404:
		if isDevMode {
			response.context.JSON(status, gin.H{"error": err.Error()})
			return
		}

		response.context.JSON(status, gin.H{"error": "Resource not found"})
		return
	case 500:
		if isDevMode {
			response.context.JSON(status, gin.H{"error": err.Error()})
			return
		}

		response.context.JSON(status, gin.H{"error": "Internal server error"})
		return
	}

}

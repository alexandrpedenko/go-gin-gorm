package app_errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHTTPError(ctx *gin.Context, err error) {
	var statusCode int
	var message string

	switch e := err.(type) {
	case *HttpError:
		statusCode = e.Code
		message = e.Message

	default:
		statusCode = http.StatusInternalServerError
		message = "An unexpected error occurred"
	}

	ctx.JSON(statusCode, gin.H{"error": message})
}

package exception

import (
	"github.com/gin-gonic/gin"
)

func ApiError(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ApiMessage(msg string, status int) gin.H {
	return gin.H{
		"status": status,
		"message": msg,
	}
}
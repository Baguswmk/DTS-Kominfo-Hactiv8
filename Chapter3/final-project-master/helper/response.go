package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestResponse(c *gin.Context, payload interface{}) {
	WriteJsonResponse(c, http.StatusBadRequest, gin.H{
		"error":   true,
		"message": payload,
	})
}

func InternalServerJsonResponse(c *gin.Context, payload interface{}) {
	WriteJsonResponse(c, http.StatusInternalServerError, gin.H{
		"error":   true,
		"message": payload,
	})
}

func NotFoundResponse(c *gin.Context, payload interface{}) {
	WriteJsonResponse(c, http.StatusNotFound, gin.H{
		"error":   true,
		"message": payload,
	})
}

func WriteJsonResponse(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

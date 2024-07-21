package utils

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (e *ErrorHandler) InternalServerError(c *gin.Context, err error) {
	if os.Getenv("ENV") == "development" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}

func (e *ErrorHandler) BadRequest(c *gin.Context, err error) {
	if os.Getenv("ENV") == "development" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": err.Error()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}
}

func (e *ErrorHandler) Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
}

func (e *ErrorHandler) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

func (e *ErrorHandler) Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}

func (e *ErrorHandler) Conflict(c *gin.Context, err error) {
	if os.Getenv("ENV") == "development" {
		c.JSON(http.StatusConflict, gin.H{"error": "Conflict", "details": err.Error()})
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "Conflict"})
	}
}

func (e *ErrorHandler) UnprocessableEntity(c *gin.Context, err error) {
	if os.Getenv("ENV") == "development" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unprocessable Entity", "details": err.Error()})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unprocessable Entity"})
	}
}

func (e *ErrorHandler) MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
}

func (e *ErrorHandler) NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not Implemented"})
}

func (e *ErrorHandler) ServiceUnavailable(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service Unavailable"})
}

func (e *ErrorHandler) GatewayTimeout(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Gateway Timeout"})
}

func (e *ErrorHandler) TooManyRequests(c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
}

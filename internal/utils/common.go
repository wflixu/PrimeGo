package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a helper function to standardize API responses.
func Response(c *gin.Context, data interface{}, err error, code ...int) {
	if err != nil {
		statusCode := http.StatusInternalServerError
		if len(code) > 0 {
			statusCode = code[0]
		}
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
		"msg":  "success",
	})
}

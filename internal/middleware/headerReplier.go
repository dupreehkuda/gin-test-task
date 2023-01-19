package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// HeaderReplier checks specific header and replies if exist
func (m middleware) HeaderReplier(c *gin.Context) {
	checkString := c.GetHeader("X-PING")

	if strings.ToLower(checkString) == "ping" {
		c.Header("X-PONG", "pong")
	}

	c.Next()
}

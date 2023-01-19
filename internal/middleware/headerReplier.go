package middleware

import "github.com/gin-gonic/gin"

// HeaderReplier checks specific header and replies if exist
func (m middleware) HeaderReplier(c *gin.Context) {
	c.Next()
}

package server

import "github.com/gin-gonic/gin"

// GetRouter returns configured gin router
func (s server) GetRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(s.mw.HeaderReplier)

	r.GET("api/when/:year", s.handlers.TimeMeasurer())

	return r
}

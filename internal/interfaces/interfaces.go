package interfaces

import "github.com/gin-gonic/gin"

type Middleware interface {
	HeaderReplier(c *gin.Context)
}

type Handlers interface {
	TimeMeasurer() gin.HandlerFunc
}

type Processors interface {
	TimeMeasurer(year int) (int, error)
}

package interfaces

import (
	"github.com/gin-gonic/gin"

	"github.com/dupreehkuda/gin-test-task/internal/models"
)

type Middleware interface {
	HeaderReplier(c *gin.Context)
}

type Handlers interface {
	TimeMeasurer() gin.HandlerFunc
}

type Processors interface {
	TimeMeasurer(year int) *models.DaysSubYearResponse
}

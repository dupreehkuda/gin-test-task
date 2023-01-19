package processors

import (
	"time"

	"github.com/dupreehkuda/gin-test-task/internal/models"
)

// TimeMeasurer measures how many days since/till specific year
func (p processor) TimeMeasurer(year int) *models.DaysSubYearResponse {
	requested := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	timeInDays := int(time.Until(requested).Hours()) / 24

	return &models.DaysSubYearResponse{
		RequestedYear: year,
		DaysToYear:    timeInDays,
	}
}

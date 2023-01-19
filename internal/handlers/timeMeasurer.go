package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TimeMeasurer handles the operation of measuring time since/till specific year
func (h handlers) TimeMeasurer() gin.HandlerFunc {
	return func(c *gin.Context) {
		yearString := c.Param("year")

		if yearString == "" {
			h.logger.Error().Msg("Wrong format")
			c.Status(http.StatusBadRequest)
			return
		}

		year, err := strconv.Atoi(yearString)
		if err != nil {
			h.logger.Err(err).Msg("Error occurred while converting year")
			c.Status(http.StatusInternalServerError)
			return
		}

		days := h.processor.TimeMeasurer(year)
		
		c.JSON(http.StatusOK, days)
	}
}

package models

// DaysSubYearResponse provides response for 'when' request
type DaysSubYearResponse struct {
	RequestedYear int `json:"requestedYear"`
	DaysToYear    int `json:"daysToYear"`
}

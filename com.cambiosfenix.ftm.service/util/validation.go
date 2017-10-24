package util

import (
	"time"

	"FenixPlatform/com.cambiosfenix.ftm.service/log"
)

// If the dates are valid return true, start date minor that end date
func ValidateDates(pStartDate *string, pIsCreate bool) (bool, string) {
	var message string
	//var startDate time.Time
	//var err error

	if pStartDate == nil {
		startDateTmp := ""
		pStartDate = &startDateTmp
	}

	if pIsCreate {
		if *pStartDate == "" {
			*pStartDate = time.Now().Format("2006-01-02")
		} else if *pStartDate != "" {
			startDate, err := time.Parse("2006-01-02", *pStartDate)
			_ = startDate
			if err != nil {
				message = "The start date is incorrect, format should be YYYY-MM-DD."
				log.Error(message)
				return false, message
			}

		}
	}
	return true, ""
}

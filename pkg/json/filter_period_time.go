package json

import (
	"time"

	"github.com/tuanden0/tx-report/pkg/consts"
)

func filterPeriodTime(periodTime, rDate string) (bool, error) {
	// Parse Date to filter
	// User time.Parse to handle data must be correct
	date, err := time.Parse(consts.DateFormat, rDate)
	if err != nil {
		return false, err
	}
	// Filter date first to avoid parsing another unused fields
	if date.Format(consts.PeriodTimeFormat) != periodTime {
		return false, nil
	}

	return true, nil
}

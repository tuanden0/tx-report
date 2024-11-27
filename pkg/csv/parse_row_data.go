package csv

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/tuanden0/tx-report/pkg/consts"
	"github.com/tuanden0/tx-report/pkg/model"
)

func parseRowDataWithPeriodTimeFilter(
	periodTime, rDate, rAmount, rContent string,
) (*model.Row, error) {
	// Parse Date to filter
	// User time.Parse to handle data must be correct
	date, err := time.Parse(consts.DateFormat, rDate)
	if err != nil {
		return nil, err
	}
	// Filter date first to avoid parsing another unused fields
	if date.Format(consts.PeriodTimeFormat) != periodTime {
		return nil, nil
	}

	// Parse amount
	amount, err := decimal.NewFromString(rAmount)
	if err != nil {
		return nil, err
	}

	return &model.Row{
		Date:    rDate,
		Amount:  amount,
		Content: rContent,
	}, nil
}

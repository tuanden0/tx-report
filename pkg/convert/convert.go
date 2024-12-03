package convert

import (
	"encoding/json"

	"github.com/shopspring/decimal"
	"github.com/tuanden0/tx-report/pkg/model"
)

func ToReply(periodTime string, rows []*model.Row) *model.Reply {
	var (
		totalIncome      = decimal.NewFromInt(0)
		totalExpenditure = decimal.NewFromInt(0)
		nRows            = len(rows)
	)

	// Calculate total data
	for i := 0; i < nRows; i++ {
		var amount = rows[i].Amount

		// Calculate total expenditure
		if amount.IsNegative() {
			totalExpenditure = totalExpenditure.Add(amount)

			continue
		}

		// Calculate total income
		totalIncome = totalIncome.Add(amount)
	}

	var (
		tic, _ = totalIncome.Float64()
		te, _  = totalExpenditure.Float64()
	)

	return &model.Reply{
		Period:           periodTime,
		TotalIncome:      tic,
		TotalExpenditure: te,
		Transactions:     rows,
	}
}

func ToJSONString(periodTime string, rows []*model.Row) (string, error) {
	data := ToReply(periodTime, rows)
	if data == nil {
		return "", nil
	}

	s, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return "", err
	}

	return string(s), nil
}

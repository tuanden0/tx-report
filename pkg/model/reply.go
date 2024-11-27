package model

type Reply struct {
	Period           string  `json:"period"`
	TotalIncome      float64 `json:"total_income"`      // Must be string to avoid rounding if amount has decimal number
	TotalExpenditure float64 `json:"total_expenditure"` // Must be string to avoid rounding if amount has decimal number
	Transactions     []*Row  `json:"transactions"`
}

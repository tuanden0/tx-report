package model

import (
	"github.com/shopspring/decimal"
)

type Row struct {
	Date    string          `json:"date"`
	Amount  decimal.Decimal `json:"amount"` // Using decimal.Decimal to avoid floating number calculation
	Content string          `json:"content"`
}

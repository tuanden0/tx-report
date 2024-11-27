package validate

import "fmt"

func RowRequiredData(date, amount, content string) error {
	if date == "" {
		return fmt.Errorf("date is empty")
	}
	if amount == "" {
		return fmt.Errorf("amount is empty")
	}
	if content == "" {
		return fmt.Errorf("content is empty")
	}

	return nil
}

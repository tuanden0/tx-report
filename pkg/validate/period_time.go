package validate

import (
	"fmt"
	"time"

	"github.com/tuanden0/tx-report/pkg/consts"
)

func PeriodTime(input string) error {
	// Input must be required
	if input == "" {
		return fmt.Errorf("yearAndMonth input is empty")
	}

	// Validate YYYYMM format
	if _, err := time.Parse(consts.PeriodTimeFormat, input); err != nil {
		return fmt.Errorf("yearAndMonth input is invalid due to: %w", err)
	}

	return nil
}

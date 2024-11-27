package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/tuanden0/tx-report/pkg/model"
	"github.com/tuanden0/tx-report/pkg/validate"
)

func Read(periodTime, f string) ([]*model.Row, error) {
	// Open CSV file
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("filePath[%q] cannot be open due to: %w", f, err)
	}
	defer file.Close()

	// Read CSV file content
	var (
		reader = csv.NewReader(file)
		rowIdx = 0
		rows   []*model.Row
	)

	// Read each line to save a memory
	for {
		rowIdx++
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("row[%d] cannot be read due to: %w", rowIdx, err)
		}

		// Validate CSV row valid
		if len(row) != csvHeaderLength {
			return nil, fmt.Errorf("row[%d] content has invalid data: %v", rowIdx, row)
		}

		// Handle CSV Header line
		if rowIdx == 1 && isHeader(row) {
			continue
		}

		var (
			rDate    = row[dateIdx]
			rAmount  = row[amountIdx]
			rContent = row[contentIdx]
		)

		// Validate row data
		if err = validate.RowRequiredData(rDate, rAmount, rContent); err != nil {
			return nil, fmt.Errorf("row[%d] content has invalid data due to: %w", rowIdx, err)
		}

		// Parse row data
		rowData, err := parseRowDataWithPeriodTimeFilter(periodTime, rDate, rAmount, rContent)
		if err != nil {
			return nil, fmt.Errorf("row[%d] content has invalid data due to: %w", rowIdx, err)
		}
		if rowData == nil {
			continue
		}

		rows = append(rows, rowData)
	}

	return rows, nil
}

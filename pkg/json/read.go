package json

import (
	"encoding/json"
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

	// Read JSON file content
	var (
		decoder = json.NewDecoder(file)
		rowIdx  = 0
		rows    []*model.Row
	)

	// read open bracket
	_, err = decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("filePath[%q] cannot be read due to: %w", f, err)
	}

	// Read each line to save a memory
	for decoder.More() {
		var (
			row    *model.Row
			rowIdx = rowIdx + 1
		)

		if err := decoder.Decode(&row); err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("row[%d] cannot be read due to: %w", rowIdx, err)
		}
		if row == nil {
			continue
		}

		// Validate row data
		var rDate = row.Date
		if err = validate.RowRequiredData(
			rDate, row.Amount.String(), row.Content,
		); err != nil {
			return nil, fmt.Errorf("row[%d] content has invalid data due to: %w", rowIdx, err)
		}

		// Filter period time
		isValid, err := filterPeriodTime(periodTime, rDate)
		if err != nil {
			return nil, fmt.Errorf("row[%d] content has invalid data due to: %w", rowIdx, err)
		}
		if !isValid {
			continue
		}

		rows = append(rows, row)
	}

	// read closing bracket
	_, err = decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("filePath[%q] cannot be read due to: %w", f, err)
	}

	return rows, nil
}

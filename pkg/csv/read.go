package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/tuanden0/tx-report/pkg/model"
	"github.com/tuanden0/tx-report/pkg/validate"
)

func Read(periodTime, f string) ([]*model.Row, error) {
	// Open CSV file
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("filePath[%q] cannot be opened due to: %w", f, err)
	}
	defer file.Close()

	// Channels for rows, results, and errors
	// Worker pools to process data
	var (
		workerCount = runtime.NumCPU() // Only work for local machine
		rowChan     = make(chan []string)
		resultChan  = make(chan *model.Row)
		errChan     = make(chan error, 1)
		wg          sync.WaitGroup
	)

	// Goroutine to read rows from CSV
	go func() {
		defer close(rowChan)

		var (
			rowIdx = 0
			reader = csv.NewReader(file) // Initialize CSV reader
		)

		for {
			rowIdx++
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				errChan <- fmt.Errorf("row[%d] cannot be read due to: %w", rowIdx, err)
				return
			}

			// Skip header
			if rowIdx == 1 && isHeader(row) {
				continue
			}

			rowChan <- row
		}
	}()

	// Worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowChan {
				// Validate CSV row structure
				if len(row) != csvHeaderLength {
					errChan <- fmt.Errorf("invalid row length: %v", row)
					return
				}

				// Extract fields
				rDate, rAmount, rContent := row[dateIdx], row[amountIdx], row[contentIdx]

				// Validate required data
				if err := validate.RowRequiredData(rDate, rAmount, rContent); err != nil {
					errChan <- err
					return
				}

				// Parse row data
				rowData, err := parseRowDataWithPeriodTimeFilter(periodTime, rDate, rAmount, rContent)
				if err != nil {
					errChan <- err
					return
				}
				if rowData != nil {
					resultChan <- rowData
				}
			}
		}()
	}

	// Goroutine to wait for workers and close resultChan
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan) // Close error channel after processing
	}()

	// Collect results and handle errors
	var rows []*model.Row
	for {
		select {
		case rowData, ok := <-resultChan:
			if ok {
				rows = append(rows, rowData)
			} else {
				resultChan = nil
			}
		case err, ok := <-errChan:
			if ok {
				return nil, err // Return immediately on error
			} else {
				errChan = nil
			}
		}

		if resultChan == nil && errChan == nil {
			break // Exit when all channels are closed
		}
	}

	return rows, nil
}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tuanden0/tx-report/pkg/consts"
	"github.com/tuanden0/tx-report/pkg/convert"
	"github.com/tuanden0/tx-report/pkg/csv"
	"github.com/tuanden0/tx-report/pkg/json"
	"github.com/tuanden0/tx-report/pkg/model"
	"github.com/tuanden0/tx-report/pkg/validate"
)

func main() {
	// Parse user input args
	var (
		periodTime string
		filePath   string
	)

	flag.StringVar(&periodTime, "p", "", "Period time in YYYYMM format (required).")
	flag.StringVar(&filePath, "f", "", "Path to csv or json file containing reports (required).")
	flag.Parse()

	// Validate user input
	if err := validate.PeriodTime(periodTime); err != nil {
		log.Fatalln(err)
	}
	fExt, err := validate.FilePath(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	var rows []*model.Row

	switch fExt {
	case consts.ExtCSV:
		// Read CSV file
		rows, err = csv.Read(periodTime, filePath)
		if err != nil {
			log.Fatalln(err)
		}
	case consts.ExtJSON:
		rows, err = json.Read(periodTime, filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Convert to JSON string
	result, err := convert.ToJSONString(periodTime, rows)
	if err != nil {
		log.Fatalf("filePath[%q] cannot convert to output data due to: %v\n", filePath, err)
	}

	// Print output
	fmt.Println(result)
}

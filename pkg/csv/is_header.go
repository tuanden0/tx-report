package csv

import "strings"

func isHeader(row []string) bool {
	return strings.ToLower(strings.Join(row, `,`)) == csvHeaderLine
}

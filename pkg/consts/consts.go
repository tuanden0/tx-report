package consts

const (
	PeriodTimeFormat = `200601`
	DateFormat       = `2006/01/02`
	ExtJSON          = ".json"
	ExtCSV           = ".csv"
)

var validExtFile = map[string]bool{
	ExtCSV:  true,
	ExtJSON: true,
}

func GetValidExtFile(ext string) bool {
	return validExtFile[ext]
}

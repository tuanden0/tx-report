package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tuanden0/tx-report/pkg/consts"
)

func FilePath(f string) (string, error) {
	// Input `f` must be required
	if f == "" {
		return "", fmt.Errorf("filePath input is empty")
	}

	// Check `f` exists
	fileInfo, err := os.Stat(f)
	if err != nil {
		return "", fmt.Errorf("unable to read filePath[%q] input due to: %w", f, err)
	}

	// Check `f` is file
	if fileInfo.IsDir() {
		return "", fmt.Errorf("filePath[%q] input is not a file", f)
	}

	// Check `f` extension
	var fExt = filepath.Ext(f)
	if !consts.GetValidExtFile(fExt) {
		return "", fmt.Errorf("filePath[%q] input must be '.csv' or '.json' file", f)
	}

	return strings.ToLower(fExt), nil
}

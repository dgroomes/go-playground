package

import (
	"os"
)

type FileSummary struct {
	Name string
	Size int64
}

// SummarizeFiles lists files in the current working directory and describes each file with a brief summary.
func SummarizeFiles() ([]FileSummary, error) {
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(currentWorkingDir)
	if err != nil {
		return nil, err
	}

	fileNames, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	fileSummaries := make([]FileSummary, len(fileNames))
	for idx, name := range fileNames {
		fi, err := os.Stat(name)
		if err != nil {
			return nil, err
		}
		size := fi.Size()
		fileSummaries[idx] = FileSummary{Name: name, Size: size}
	}

	return fileSummaries, nil
}

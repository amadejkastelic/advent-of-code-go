package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amadejkastelic/advent-of-code-go/pkg/aoc"
)

const (
	cachePath = "./inputs/%d/day_%d/%s.txt"

	simpleInputFileName = "simple_input"
	actualInputFileName = "input"
)

// FetchInput retrieves the puzzle input path for the specified year and day.
func FetchInput(year int, day int, simple bool) (string, error) {
	filePath := fmt.Sprintf(cachePath, year, day, fileName(simple))

	if fileInfo, err := os.Stat(filePath); err == nil && !fileInfo.IsDir() {
		return filePath, nil
	}

	client, err := aoc.NewAOCClient()
	if err != nil {
		return "", err
	}

	inp, err := client.FetchInput(year, day, simple)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	_, err = f.WriteString(inp)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func fileName(simple bool) string {
	if simple {
		return simpleInputFileName
	}
	return actualInputFileName
}

package common

import (
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromInput(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(bytes))
	lines := strings.Split(trimmed, "\n")

	return lines, nil
}

func Int(n string) int {
	i, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	return i
}

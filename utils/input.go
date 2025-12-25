package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	AOC_URL        = "https://adventofcode.com/2025/day/%d/input"
	INPUT_FILENAME = "input.txt"
)

func GetInput(day int) (io.ReadCloser, error) {
	_, callingFileName, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("Couldn't get filepath")
	}
	dir := filepath.Dir(callingFileName)
	fullInputFilePath := filepath.Join(dir, INPUT_FILENAME)
	file, err := os.Open(fullInputFilePath)
	if err == nil {
		return file, nil
	}
	file, err = os.Create(fullInputFilePath)
	if err != nil {
		return nil, err
	}
	err = downloadInput(fmt.Sprintf(AOC_URL, day), file)
	file.Close()
	if err != nil {
		return nil, err
	}
	file, err = os.Open(fullInputFilePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

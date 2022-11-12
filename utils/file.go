package utils

import (
	"os"

	"github.com/aosasona/stripr/types"
)

func CheckFileExists(path string) bool {
	if fileExists, err := os.Stat(path); err != nil || fileExists.IsDir() {
		return false
	}
	return true
}

func CheckDirExists(path string) bool {
	if dirExists, err := os.Stat(path); err != nil || !dirExists.IsDir() {
		return false
	}
	return true
}

func ReadFileAsString(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", &types.CustomError{Message: "Unable to read file"}
	}
	return string(file), nil
}

func ReadDirectory(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, &types.CustomError{Message: "Unable to read directory"}
	}
	return files, nil
}

package utils

import "os"

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

func ReadFileAsString(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		Terminate(err)
	}
	return string(file)
}

func ReadDirectory(path string) []os.DirEntry {
	files, err := os.ReadDir(path)
	if err != nil {
		Terminate(err)
	}
	return files
}

package utils

import (
	"github.com/aosasona/stripr/types"
	"os"
)

type Scanner types.Scanner

func (s *Scanner) New(path string) {
	s.Path = path
	if s.DirExists() {
		s.DirType = types.DIRECTORY
	} else {
		s.DirType = types.FILE
	}
}

func (s *Scanner) ScanSingle(path types.FilePath) (types.ScanResult, error) {
	if !s.FileExists(path) {
		return types.ScanResult{}, &types.ErrFileNotFound{Path: string(path)}
	}

	return types.ScanResult{}, nil
}

func (s *Scanner) ScanDir(path types.FilePath) ([]types.ScanResult, error) {
	return []types.ScanResult{}, nil
}

func (s *Scanner) FileExists(path types.FilePath) bool {
	if fileExists, err := os.Stat(string(path)); err != nil || fileExists.IsDir() {
		return false
	}
	return true
}

func (s *Scanner) DirExists() bool {
	if dirExists, err := os.Stat(string(s.Path)); err != nil || !dirExists.IsDir() {
		return false
	}
	return true
}

func (s *Scanner) CountDirFiles() int {
	path := s.Path
	if !s.DirExists() {
		return 0
	}

	count := 0

	files, err := os.ReadDir(string(path))
	if err != nil {
		return 0
	}

	for _, file := range files {
		if !file.IsDir() && file.Type().IsRegular() {
			count++
		}
	}

	return count
}

func (s *Scanner) IsTextFile(path types.FilePath) bool {

	if !s.FileExists(path) {
		return false
	}

	return true
}

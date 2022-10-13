package stripr

import (
	"github.com/aosasona/stripr/types"
	"github.com/aosasona/stripr/utils"
	"os"
)

type Scanner types.Scanner

func (s *Scanner) New() *Scanner {
	if utils.CheckDirExists(s.Path) {
		s.DirType = types.DIRECTORY
		return s
	}
	s.DirType = types.FILE
	return s
}

func (s *Scanner) Scan() interface{} {
	switch s.DirType {
	case types.FILE:
		file, err := s.ScanSingle()
		if err != nil {
			utils.Terminate(&types.FatalRuntimeError{})
		}
		return &file
	case types.DIRECTORY:
		files, err := s.ScanDir()
		if err != nil {
			utils.Terminate(&types.FatalRuntimeError{})
		}
		return &files
	default:
		utils.Terminate(&types.FatalRuntimeError{})
		break
	}
	return nil
}

func (s *Scanner) ScanSingle() (types.ScanResult, error) {
	if !utils.CheckFileExists(s.Path) {
		return types.ScanResult{}, &types.ErrFileNotFound{Path: string(s.Path)}
	}

	return types.ScanResult{}, nil
}

func (s *Scanner) ScanDir() ([]types.ScanResult, error) {
	if !utils.CheckDirExists(s.Path) {
		return nil, &types.ErrDirNotFound{Path: s.Path}
	}

	return []types.ScanResult{}, nil
}

func (s *Scanner) CountDirFiles() (int, error) {
	path := s.Path
	if !utils.CheckDirExists(s.Path) {
		return 0, &types.ErrDirNotFound{Path: path}
	}

	count := 0

	files, err := os.ReadDir(string(path))
	if err != nil {
		return 0, &types.UnableToReadDir{Path: path}
	}

	for _, file := range files {
		if !file.IsDir() && file.Type().IsRegular() {
			count++
		}
	}

	return count, nil
}

func (s *Scanner) IsTextFile(path types.FilePath) bool {

	return true
}

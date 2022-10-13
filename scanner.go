package main

import (
	"github.com/aosasona/stripr/types"
	"github.com/aosasona/stripr/utils"
)

type Scanner types.Scanner

func (s *Scanner) New(args []string) *Scanner {

	if len(args) < 2 {
		utils.Terminate(&types.ErrNoCommand{})
	}

	s.Args = args
	s.Path = args[1]

	if utils.CheckDirExists(s.Path) {
		s.DirType = types.DIRECTORY
		return s
	}
	s.DirType = types.FILE
	return s
}

func (s *Scanner) Run() {
	commands := []string{"scan", "clean", "init"}
	args := s.Args

	if len(args) > 2 {
		if !utils.Contains(commands, args[2]) {
			utils.Terminate(&types.ErrInvalidCommand{Command: args[0]})
		}
	}
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
	count := len(utils.ReadDirectory(path))
	return count, nil
}

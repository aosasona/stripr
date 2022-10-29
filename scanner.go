package main

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/aosasona/stripr/types"
	"github.com/aosasona/stripr/utils"
)

type Scanner types.Scanner

const (
	CommentRegex = `(/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/)|(//.*)`
)

func (s *Scanner) New(dirPath *string) *Scanner {

	s.Path = *dirPath

	if utils.CheckDirExists(s.Path) {
		s.DirType = types.DIRECTORY
	} else if utils.CheckFileExists(s.Path) {
		s.DirType = types.FILE
		return s
	} else {
		log.Fatalf("Path %s does not exist", s.Path)
	}

	err := s.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	return s
}

func (s *Scanner) Init() error {

	if s.DirType == types.FILE {
		return errors.New("cannot create config file in a file")
	}

	configContent := []byte(`{
	"ignore": ["node_modules", "tests", "vendor", "dist", "build"],
	"showStats": true,
	"confirmStrip": true
}`)

	path := strings.Trim(s.Path, "/") + "/stripr.json"
	err := os.WriteFile(path, configContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Scanner) Scan() ([]map[string]interface{}, int, error) {
	switch s.DirType {
	case types.FILE:
		file, _, err := s.ScanSingle()
		if err != nil {
			return nil, 0, err
		}
		return []map[string]interface{}{utils.StructToMap(file)}, 0, nil
	case types.DIRECTORY:
		files, ignoredCount, err := s.ScanDir()
		if err != nil {
			return nil, 0, err
		}
		return utils.StructsToMaps(files), ignoredCount, nil
	default:
		utils.Terminate(&types.FatalRuntimeError{})
		break
	}
	return nil, 0, nil
}

func (s *Scanner) ScanSingle() (types.ScanResult, int, error) {
	if !utils.CheckFileExists(s.Path) {
		return types.ScanResult{}, 0, &types.ErrFileNotFound{Path: string(s.Path)}
	}

	comments := s.GetComments(s.Path)
	hasComments := len(comments) > 0

	scanResult := types.ScanResult{
		Name:        s.Path,
		Path:        s.Path,
		Lines:       comments,
		HasComments: hasComments,
	}

	return scanResult, 0, nil
}

func (s *Scanner) ScanDir() ([]types.ScanResult, int, error) {
	if !utils.CheckDirExists(s.Path) {
		return nil, 0, &types.ErrDirNotFound{Path: s.Path}
	}

	files := utils.ReadDirectory(s.Path)
	var scanResults []types.ScanResult
	ignoredCount := 0

	for _, file := range files {

		ignored := s.CheckIfFileIgnored(file.Name())
		if ignored {
			ignoredCount++
			continue
		}

		comments := s.GetComments(s.Path + "/" + file.Name())
		hasComments := len(comments) > 0

		scanResult := types.ScanResult{
			Name:        file.Name(),
			Path:        s.Path + "/" + file.Name(),
			Lines:       comments,
			HasComments: hasComments,
		}
		scanResults = append(scanResults, scanResult)
	}

	return scanResults, ignoredCount, nil
}

func (s *Scanner) CountDirFiles() (int, error) {
	path := s.Path
	if !utils.CheckDirExists(s.Path) {
		return 0, &types.ErrDirNotFound{Path: path}
	}
	count := len(utils.ReadDirectory(path))
	return count, nil
}

func (s *Scanner) CheckIfFileIgnored(path string) bool {
	for _, ignore := range s.Config["ignore"].([]interface{}) {
		if strings.Contains(path, ignore.(string)) {
			return true
		}
	}
	return false
}

func (s *Scanner) LoadConfig() error {
	config := utils.ReadConfig(s.Path)
	s.Config = config
	return nil
}

func (s *Scanner) GetComments(file string) [][]int {
	commentRegex := regexp.MustCompile(CommentRegex)
	fileContent := utils.ReadFileAsString(file)
	matches := commentRegex.FindAllStringIndex(fileContent, -1)
	return matches
}

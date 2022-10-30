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
	"skipCheck": true
}`)

	path := strings.Trim(s.Path, "/") + "/stripr.json"
	err := os.WriteFile(path, configContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Scanner) Scan() ([]types.ScanResult, int, error) {
	var (
		scanResults  []types.ScanResult
		ignoredCount int
	)
	switch s.DirType {
	case types.FILE:
		file, _, err := s.ScanSingle()
		if err != nil {
			return []types.ScanResult{}, 0, err
		}
		ignoredCount = 0
		scanResults = append(scanResults, file)
		break
	case types.DIRECTORY:
		files, count, err := s.ScanDir()
		if err != nil {
			return []types.ScanResult{}, 0, err
		}
		ignoredCount = count
		scanResults = files
		break
	default:
		utils.Terminate(&types.FatalRuntimeError{})
		break
	}
	return scanResults, ignoredCount, nil
}

func (s *Scanner) ScanSingle() (types.ScanResult, int, error) {
	if !utils.CheckFileExists(s.Path) {
		return types.ScanResult{}, 0, &types.ErrFileNotFound{Path: string(s.Path)}
	}

	comments := s.GetComments(s.Path)

	splitName := strings.Split(s.Path, "/")
	fileName := splitName[len(splitName)-1]

	scanResult := types.ScanResult{
		Name:        fileName,
		Path:        s.Path,
		Lines:       comments,
		HasComments: len(comments) > 0,
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

		if file.IsDir() {
			continue
		}

		comments := s.GetComments(s.Path + "/" + file.Name())

		scanResult := types.ScanResult{
			Name:        file.Name(),
			Path:        s.Path + "/" + file.Name(),
			Lines:       comments,
			HasComments: len(comments) > 0,
		}
		scanResults = append(scanResults, scanResult)
	}

	return scanResults, ignoredCount, nil
}

func (s *Scanner) StripComments(name string) error {
	var filePath string
	if s.DirType == types.FILE {
		filePath = s.Path
	} else {
		filePath = s.Path + "/" + name
	}
	if !utils.CheckFileExists(filePath) {
		return &types.ErrFileNotFound{Path: filePath}
	}

	initialFileContent := utils.ReadFileAsString(filePath)
	fileContent := regexp.MustCompile(CommentRegex).ReplaceAllString(initialFileContent, "")
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
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
	var customConfig []interface{}
	if s.Config != nil {
		customConfig = s.Config["ignore"].([]interface{})
	}
	ignoredPaths := append(
		[]interface{}{".git", ".DS_Store", ".idea", ".vscode", "stripr.json", "node_modules", "tests", "vendor", "yarn.lock", "package-lock.json"},
		customConfig...,
	)

	ignoredPaths = utils.RemoveDuplicates(ignoredPaths)

	for _, ignore := range ignoredPaths {
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

package main

import (
	"fmt"
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

func (s *Scanner) New(dirPath *string) (*Scanner, error) {

	s.Path = *dirPath

	if utils.CheckDirExists(s.Path) {
		s.DirType = types.DIRECTORY
	} else if utils.CheckFileExists(s.Path) {
		s.DirType = types.FILE
		return s, nil
	} else {
		return nil, &types.CustomError{Message: fmt.Sprintf("Path %s does not exist", s.Path)}
	}

	err := s.LoadConfig()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Scanner) Init() error {

	if s.DirType == types.FILE {
		return &types.CustomError{Message: fmt.Sprintf("Path %s is a file, not a directory", s.Path)}
	}

	configContent := []byte(`{
	"ignore": ["node_modules", "tests", "vendor", "dist", "build", ".dockerignore", ".gitignore", ".env", "yarn.lock", "package.json", "package-lock.json", "composer.json", "composer.lock", "Dockerfile"],
	"showStats": true,
	"skipCheck": true
}`)

	path := strings.Trim(s.Path, "/") + "/stripr.json"
	err := os.WriteFile(path, configContent, 0644)
	if err != nil {
		return &types.CustomError{Message: fmt.Sprintf("Could not create config file at %s", path)}
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
		return []types.ScanResult{}, 0, &types.CustomError{Message: "Invalid directory type"}
	}
	return scanResults, ignoredCount, nil
}

func (s *Scanner) ScanSingle() (types.ScanResult, int, error) {
	if !utils.CheckFileExists(s.Path) {
		return types.ScanResult{}, 0, &types.CustomError{Message: fmt.Sprintf("File %s does not exist", s.Path)}
	}

	comments, err := s.GetComments(s.Path)
	if err != nil {
		return types.ScanResult{}, 0, err
	}

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
		return nil, 0, &types.CustomError{Message: fmt.Sprintf("Path %s does not exist", s.Path)}
	}

	files, err := utils.ReadDirectory(s.Path)

	if err != nil {
		return nil, 0, err
	}

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

		comments, err := s.GetComments(s.Path + "/" + file.Name())

		if err != nil {
			return nil, 0, err
		}

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
		return &types.CustomError{Message: fmt.Sprintf("File %s does not exist", filePath)}
	}

	initialFileContent, err := utils.ReadFileAsString(filePath)
	if err != nil {
		return err
	}
	fileContent := regexp.MustCompile(CommentRegex).ReplaceAllString(initialFileContent, "")
	err = os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Scanner) CountDirFiles() (int, error) {
	path := s.Path
	if !utils.CheckDirExists(s.Path) {
		return 0, &types.CustomError{Message: fmt.Sprintf("Path %s does not exist", path)}
	}
	files, err := utils.ReadDirectory(path)
	if err != nil {
		return 0, err
	}
	count := len(files)
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
	config, err := utils.ReadConfig(s.Path)
	if err != nil {
		return err
	}
	s.Config = config
	return nil
}

func (s *Scanner) GetComments(file string) ([][]int, error) {
	commentRegex := regexp.MustCompile(CommentRegex)
	fileContent, err := utils.ReadFileAsString(file)
	if err != nil {
		return nil, err
	}
	matches := commentRegex.FindAllStringIndex(fileContent, -1)
	return matches, nil
}

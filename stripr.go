package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aosasona/stripr/utils"
)

type Stripr struct {
	Target  string
	Args    []string
	Scanner *Scanner
}

type StriprOpts struct {
	Target    *string
	ShowStats bool
	Args      []string
}

func CreateStriprInstance(opts StriprOpts) *Stripr {
	scanner := Scanner{}
	scanner.New(opts.Target)
	return &Stripr{
		Target:  *opts.Target,
		Args:    opts.Args,
		Scanner: &scanner,
	}
}

func (s *Stripr) Run() *Stripr {
	if len(s.Args) < 1 {
		s.ShowUsage()
		return s
	}

	println("Current directory:", s.Target)

	mainCmd := s.Args[0]

	switch mainCmd {
	case "init":
		s.CreateConfig()
		break
	case "scan":
		s.ScanTarget()
	default:
		s.ShowUsage()
	}

	return s
}

func (s *Stripr) ShowUsage() {
	usage := `Usage: stripr [options] [command]
		
	Options:
		-target=string
				The directory or file to read (default ".")
		-show-stats		
				Show the number of files and lines that will be affected
	Commands:
		init		
				Create a config file in the current directory
		scan		
				Scan the directory for comments
		clean		
				Remove comments from the directory (-y to prevent asking for confirmation; use with caution)
		help
				Show this help message
`
	println(usage)
}

func (s *Stripr) CreateConfig() {
	err := s.Scanner.Init()
	if err != nil {
		utils.Terminate(errors.New(fmt.Sprintf("Error creating config file: %s", err)))
	}

	fmt.Printf("Config file created at %s/stripr.json", s.Target)
}

func (s *Stripr) ScanTarget() {
	stats, ignoredCount, err := s.Scanner.Scan()
	if err != nil {
		utils.Terminate(errors.New(fmt.Sprintf("Error scanning target: %s", err)))
	}

	var filesWithComments []map[string]interface{}
	for _, file := range stats {
		if file["HasComments"].(bool) {
			filesWithComments = append(filesWithComments, file)
		}
	}

	fmt.Printf("[scan] %d file(s) scanned, %d file(s) ignored\n", len(stats), ignoredCount)
	utils.PrintStats(filesWithComments)
	os.Exit(0)
}

func (s *Stripr) CleanTarget() {
	// TODO
}

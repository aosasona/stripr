package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aosasona/stripr/types"
	"github.com/aosasona/stripr/utils"
)

type Stripr struct {
	Target    string
	Args      []string
	ShowStats bool
	SkipCheck bool
	Scanner   *Scanner
}

func CreateStriprInstance(target *string, opts Stripr) *Stripr {
	scanner := Scanner{}
	scanner.New(target)
	return &Stripr{
		Target:    *target,
		ShowStats: opts.ShowStats,
		SkipCheck: opts.SkipCheck,
		Scanner:   &scanner,
		Args:      opts.Args,
	}
}

func (s *Stripr) Run() *Stripr {
	if len(s.Args) < 1 {
		s.ShowUsage()
		return s
	}

	mainCmd := s.Args[0]

	switch mainCmd {
	case "init":
		s.CreateConfig()
		break
	case "scan":
		s.ScanTarget()
		break
	case "strip":
		s.CleanTarget()
		break
	case "help":
		s.ShowUsage()
		break
	default:
		utils.Terminate(&types.ErrInvalidCommand{Command: mainCmd})
	}

	return s
}

func (s *Stripr) ShowUsage() {
	usage := `Usage: stripr [options] [command]
		
	Options:
		-target=string
				The directory or file to read (default "." - current directory)
		-show-stats=true|false		
				Show the number of files and lines that will be affected
		-skip-check=true|false
				Skip the confirmation prompt before stripping comments
	Commands:
		init		
				Create a config file in the current directory
		scan		
				Scan the directory for comments
		strip		
				Remove comments from the directory (-skip-check to prevent asking for confirmation; use with caution)
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
		utils.Terminate(&types.ErrReadingTarget{Path: s.Target})
	}

	filesWithComments := utils.SortScanResults(stats)

	fmt.Printf("[scan] %d file(s) scanned, %d file(s) ignored\n", len(stats), ignoredCount)
	utils.PrintStats(filesWithComments)
	os.Exit(0)
}

func (s *Stripr) CleanTarget() {
	stats, ignoredCount, err := s.Scanner.Scan()
	if err != nil {
		utils.Terminate(&types.ErrReadingTarget{Path: s.Target})
	}

	filesWithComments := utils.SortScanResults(stats)
	fmt.Printf("[scan] %d file(s) scanned, %d file(s) ignored\n", len(stats), ignoredCount)

	if len(filesWithComments) == 0 {
		fmt.Println("[scan] No files contain comments")
		os.Exit(0)
	}

	configExists := s.Scanner.Config != nil

	if s.ShowStats || (configExists && s.Scanner.Config["showStats"].(bool)) {
		utils.PrintStats(filesWithComments)
	}

	var input string

	if s.SkipCheck {
		fmt.Print("[clean] Skipping confirmation prompt\n")
		input = "y"
	} else if configExists && s.Scanner.Config["skipCheck"].(bool) && (input != "") {
		fmt.Print("[clean] Skipping confirmation prompt\n")
		input = "y"
	} else {
		fmt.Printf("[clean] %d file(s) will be affected (ensure you have a way to undo changes if things go wrong before proceeding)\n", len(filesWithComments))
		fmt.Printf("[clean] Are you sure you want to continue? (y/n) ")

		_, err := fmt.Scanln(&input)
		if err != nil {
			utils.Terminate(&types.ErrInvalidInput{Input: input})
		}
	}

	if strings.ToLower(input) == "y" {
		for _, file := range filesWithComments {
			err := s.Scanner.StripComments(file.Name)
			if err != nil {
				utils.Terminate(&types.FatalRuntimeError{})
			}
		}
		fmt.Println("[clean] Done! 🎉")
		os.Exit(0)
	} else if strings.ToLower(input) == "n" {
		utils.Terminate(errors.New("stripping aborted"))
	} else {
		utils.Terminate(&types.ErrInvalidInput{Input: input})
	}

}

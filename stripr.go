package main

import (
	"fmt"
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

func CreateStriprInstance(target *string, opts Stripr) (Stripr, error) {
	scanner := Scanner{}
	s, err := scanner.New(target)
	if err != nil {
		return Stripr{}, err
	}
	return Stripr{
		Target:    *target,
		ShowStats: opts.ShowStats,
		SkipCheck: opts.SkipCheck,
		Scanner:   s,
		Args:      opts.Args,
	}, nil
}

func (s *Stripr) Run() (*Stripr, error) {
	if len(s.Args) < 1 {
		s.ShowUsage()
		return s, nil
	}

	mainCmd := s.Args[0]
	var err error

	switch mainCmd {
	case "init":
		err = s.CreateConfig()
		break
	case "scan":
		err = s.ScanTarget()
		break
	case "strip":
		err = s.CleanTarget()
		break
	case "help":
		s.ShowUsage()
		break
	default:
		return s, &types.CustomError{Message: fmt.Sprintf("Unknown command: %s", mainCmd)}
	}

	if err != nil {
		return s, err
	}

	return s, nil
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

	You can optionally using a config file to set some options on a per-project basis.
`
	println(usage)
}

func (s *Stripr) CreateConfig() error {
	err := s.Scanner.Init()
	if err != nil {
		return err
	}

	fmt.Printf("Config file created at %s/stripr.json", s.Target)
	return nil
}

func (s *Stripr) ScanTarget() error {
	stats, ignoredCount, err := s.Scanner.Scan()
	if err != nil {
		return err
	}

	filesWithComments := utils.SortScanResults(stats)

	fmt.Printf("[scan] %d file(s) scanned, %d file(s) ignored\n", len(stats), ignoredCount)
	utils.PrintStats(filesWithComments)
	return nil
}

func (s *Stripr) CleanTarget() error {
	stats, ignoredCount, err := s.Scanner.Scan()
	if err != nil {
		return err
	}

	filesWithComments := utils.SortScanResults(stats)
	fmt.Printf("[scan] %d file(s) scanned, %d file(s) ignored\n", len(stats), ignoredCount)

	if len(filesWithComments) == 0 {
		fmt.Println("[scan] No files contain comments")
		return nil
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
			return &types.CustomError{Message: fmt.Sprintf("Error reading input: %s", err)}
		}
	}

	if strings.ToLower(input) == "y" {
		for _, file := range filesWithComments {
			err := s.Scanner.StripComments(file.Name)
			if err != nil {
				return err
			}
		}
		fmt.Println("[clean] Done! 🎉")
	} else if strings.ToLower(input) == "n" {
		return &types.CustomError{Message: "Aborted"}
	} else {
		return &types.CustomError{Message: "Invalid input"}
	}
	return nil
}

package main

import (
	"flag"
	"os"

	"github.com/aosasona/stripr/utils"
)

func main() {
	targetPath := flag.String("target", ".", "The directory or file to read")
	showStats := flag.Bool("show-stats", false, "Show the number of files and lines that will be affected")
	skipCheck := flag.Bool("skip-check", false, "Skip the confirmation prompt before stripping comments")

	flag.Parse()

	stripr, err := new(Stripr).New(targetPath, Stripr{
		ShowStats: *showStats,
		SkipCheck: *skipCheck,
		Args:      flag.Args(),
	})
	if err != nil {
		utils.FilterErrorAndTerminate(err)
	}

	_, err = stripr.Run()
	if err != nil {
		utils.FilterErrorAndTerminate(err)
	}

	os.Exit(0)
}

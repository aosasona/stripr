package main

import (
	"flag"
)

func main() {
	targetPath := flag.String("target", ".", "The directory or file to read")
	showStats := flag.Bool("show-stats", false, "Show the number of files and lines that will be affected")
	skipCheck := flag.Bool("skip-check", false, "Skip the confirmation prompt before stripping comments")

	flag.Parse()

	stripr := CreateStriprInstance(targetPath, Stripr{
		ShowStats: *showStats,
		SkipCheck: *skipCheck,
		Args:      flag.Args(),
	})
	stripr.Run()
}

package main

import (
	"flag"
)

func main() {
	targetPath := flag.String("target", ".", "The directory or file to read")
	showStats := flag.Bool("show-stats", false, "Show the number of files and lines that will be affected")

	flag.Parse()

	s := CreateStriprInstance(StriprOpts{
		Target:    targetPath,
		ShowStats: *showStats,
		Args:      flag.Args(),
	})
	s.Run()
}

package utils

import (
	"fmt"

	"github.com/aosasona/stripr/types"
)

func PrintStats(stats []types.ScanResult) {
	if len(stats) > 0 {
		fmt.Print("\n----------------------------------------\n Stats - only file(s) with comment(s) \n----------------------------------------\n")
		for _, file := range stats {
			fmt.Printf("- %s (%d lines)\n", file.Name, len(file.Lines))
		}
		fmt.Print("\n")
	} else {
		fmt.Println("[scan] No files contain comments")
	}
}

func SortScanResults(results []types.ScanResult) []types.ScanResult {
	var sorted []types.ScanResult
	for _, result := range results {
		if result.HasComments {
			sorted = append(sorted, result)
		}
	}
	return sorted
}

func RemoveDuplicates(elements []interface{}) []interface{} {
	encountered := map[string]bool{}
	var result []interface{}

	for v := range elements {
		if !encountered[elements[v].(string)] {
			encountered[elements[v].(string)] = true
			result = append(result, elements[v])
		}
	}
	return result
}

package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aosasona/stripr/types"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	b, _ := json.Marshal(s)
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return m
}

func StructsToMaps(s []types.ScanResult) []map[string]interface{} {
	var maps []map[string]interface{}
	for _, v := range s {
		maps = append(maps, StructToMap(v))
	}
	return maps
}

func PrintStats(stats []map[string]interface{}) {
	if len(stats) > 0 {
		fmt.Print("----------------------------------------\n Stats - only file(s) with comment(s) \n----------------------------------------\n")
		for _, file := range stats {
			fmt.Printf("- %s (%d lines)", file["Name"], len(file["Lines"].([]interface{})))
		}
	} else {
		fmt.Println("[scan] No files contain comments")
	}
}

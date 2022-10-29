package utils

import (
	"encoding/json"

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

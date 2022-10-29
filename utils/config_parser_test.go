package utils

import (
	"reflect"
	"testing"
)

func TestCheckConfigExists(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 bool
	}{
		{
			"check if config file exists",
			"../example/stripr.json",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckConfigExists("../example")
			if got != tt.want {
				t.Errorf("CheckConfigExists() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckConfigExists() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name string
		want map[string]interface{}
	}{
		{"attempt to read config file",
			map[string]interface{}{
				"ignore": []interface{}{
					"node_modules", "tests", "vendor", "dist", "build",
				},
				"showStats":    true,
				"confirmStrip": true,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig("../example"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

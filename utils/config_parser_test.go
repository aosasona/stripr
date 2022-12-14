package utils

import (
	"reflect"
	"testing"
)

func TestCheckConfigExists(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 interface{}
	}{
		{
			"check if config file exists",
			"../example/stripr.json",
			nil,
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
					"node_modules", "tests", "vendor", "dist", "build", ".dockerignore", ".gitignore", ".env", "yarn.lock", "package.json", "package-lock.json", "composer.json", "composer.lock", "Dockerfile",
				},
				"showStats": true,
				"skipCheck": true,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ReadConfig("../example"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

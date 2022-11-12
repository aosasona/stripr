package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestCheckDirExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"check if directory exists",
			args{"../example"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDirExists(tt.args.path); got != tt.want {
				t.Errorf("CheckDirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFileExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"check if file exists",
			args{"../example/index.js"},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFileExists(tt.args.path); got != tt.want {
				t.Errorf("CheckFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []os.DirEntry
	}{
		{
			"attempt to read a directory",
			args{"../example"},
			[]os.DirEntry{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ReadDirectory(tt.args.path); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("ReadDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

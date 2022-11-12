package stripr

import (
	"reflect"
	"testing"

	"github.com/aosasona/stripr/types"
)

func TestScanner_CheckIfFileIgnored(t *testing.T) {
	type args struct {
		path string
	}

	dirPath := "./example"
	scanner := Scanner{}
	s, _ := scanner.New(&dirPath)

	tests := []struct {
		name string
		s    Scanner
		args args
		want bool
	}{
		{
			name: "file is ignored",
			s:    *s,
			args: args{path: "./example/read.js"},
			want: true,
		},
		{
			name: "file is not ignored",
			s:    Scanner{},
			args: args{path: "./example/server.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CheckIfFileIgnored(tt.args.path); got != tt.want {
				t.Errorf("CheckIfFileIgnored() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_CountDirFiles(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CountDirFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("CountDirFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountDirFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_GetComments(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		s       Scanner
		args    args
		want    [][]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetComments(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_Init(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScanner_LoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LoadConfig(); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScanner_New(t *testing.T) {
	type args struct {
		dirPath *string
	}
	tests := []struct {
		name    string
		s       Scanner
		args    args
		want    *Scanner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.New(tt.args.dirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_Scan(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		want    []types.ScanResult
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Scan()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Scan() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScanner_ScanDir(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		want    []types.ScanResult
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.ScanDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScanDir() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ScanDir() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScanner_ScanSingle(t *testing.T) {
	tests := []struct {
		name    string
		s       Scanner
		want    types.ScanResult
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.ScanSingle()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanSingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScanSingle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ScanSingle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScanner_StripComments(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		s       Scanner
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.StripComments(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("StripComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

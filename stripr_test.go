package main

import (
	"reflect"
	"testing"
)

func TestCreateCMD(t *testing.T) {
	type args struct {
		target *string
		opts   Stripr
	}
	tests := []struct {
		name    string
		args    args
		want    *Stripr
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateStriprInstance(tt.args.target, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCMD() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripr_CleanTarget(t *testing.T) {
	type fields struct {
		Target    string
		Args      []string
		ShowStats bool
		SkipCheck bool
		Scanner   *Scanner
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stripr{
				Target:    tt.fields.Target,
				Args:      tt.fields.Args,
				ShowStats: tt.fields.ShowStats,
				SkipCheck: tt.fields.SkipCheck,
				Scanner:   tt.fields.Scanner,
			}
			if err := s.CleanTarget(); (err != nil) != tt.wantErr {
				t.Errorf("CleanTarget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStripr_CreateConfig(t *testing.T) {
	type fields struct {
		Target    string
		Args      []string
		ShowStats bool
		SkipCheck bool
		Scanner   *Scanner
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stripr{
				Target:    tt.fields.Target,
				Args:      tt.fields.Args,
				ShowStats: tt.fields.ShowStats,
				SkipCheck: tt.fields.SkipCheck,
				Scanner:   tt.fields.Scanner,
			}
			s.CreateConfig()
		})
	}
}

func TestStripr_Run(t *testing.T) {
	type fields struct {
		Target    string
		Args      []string
		ShowStats bool
		SkipCheck bool
		Scanner   *Scanner
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Stripr
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stripr{
				Target:    tt.fields.Target,
				Args:      tt.fields.Args,
				ShowStats: tt.fields.ShowStats,
				SkipCheck: tt.fields.SkipCheck,
				Scanner:   tt.fields.Scanner,
			}
			got, err := s.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripr_ScanTarget(t *testing.T) {
	type fields struct {
		Target    string
		Args      []string
		ShowStats bool
		SkipCheck bool
		Scanner   *Scanner
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stripr{
				Target:    tt.fields.Target,
				Args:      tt.fields.Args,
				ShowStats: tt.fields.ShowStats,
				SkipCheck: tt.fields.SkipCheck,
				Scanner:   tt.fields.Scanner,
			}
			if err := s.ScanTarget(); (err != nil) != tt.wantErr {
				t.Errorf("ScanTarget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStripr_ShowUsage(t *testing.T) {
	type fields struct {
		Target    string
		Args      []string
		ShowStats bool
		SkipCheck bool
		Scanner   *Scanner
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stripr{
				Target:    tt.fields.Target,
				Args:      tt.fields.Args,
				ShowStats: tt.fields.ShowStats,
				SkipCheck: tt.fields.SkipCheck,
				Scanner:   tt.fields.Scanner,
			}
			s.ShowUsage()
		})
	}
}

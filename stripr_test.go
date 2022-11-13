package main

import (
	"reflect"
	"testing"

	"github.com/aosasona/stripr/types"
)

func TestCreateCMD(t *testing.T) {
	stripr := new(Stripr)
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
		{
			name: "Test CreateCMD",
			args: args{
				target: &[]string{"./example"}[0],
				opts: Stripr{
					Args:      []string{},
					ShowStats: false,
					SkipCheck: false,
				},
			},
			want: &Stripr{
				Target:    "./example",
				Args:      []string{},
				ShowStats: false,
				SkipCheck: false,
				Scanner:   &Scanner{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stripr.New(tt.args.target, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.TypeOf(got).AssignableTo(reflect.TypeOf(tt.want)) || !reflect.TypeOf(tt.want.Scanner).AssignableTo(reflect.TypeOf(got.Scanner)) {
				t.Errorf("CreateCMD() got = %v, want %v", got, tt.want)
			}
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
		{
			name: "run stripr",
			fields: fields{
				Target:    "./example",
				Args:      []string{},
				ShowStats: false,
				SkipCheck: false,
				Scanner: &Scanner{
					Path:    "./example",
					DirType: types.DIRECTORY,
				},
			},
			want: &Stripr{
				Target:    "./example",
				Args:      []string{},
				ShowStats: false,
				SkipCheck: false,
				Scanner:   &Scanner{},
			},
			wantErr: false,
		},
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
			if !reflect.DeepEqual(got.Target, tt.want.Target) {
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
		{
			name: "scan target",
			fields: fields{
				Target:    "./example",
				Args:      []string{},
				ShowStats: false,
				SkipCheck: false,
				Scanner: &Scanner{
					Path:    "./example",
					DirType: types.DIRECTORY,
				},
			},
		},
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

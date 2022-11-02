package config

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_init_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}

func Test_crete(t *testing.T) {
	tests := []struct {
		name string
		want Content
	}{
		{
			name: "test_create_1",
			want: Content{
				Path:   "需要批量更换文件名的父目录的绝对路径",
				Expr:   "[S|s]{1}[0-9]{2}[e|E]{1}[0-9]{2}",
				Prefix: FileNamePrefix,
				Suffix: "mp4",
				Ignore: true,
				Scan:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Content
	}{
		{
			name: "test_parse_1",
			args: args{
				path: "cfn.json",
			},
			want: Content{
				Path:   "需要批量更换文件名的父目录的绝对路径",
				Expr:   "[S|s]{1}[0-9]{2}[e|E]{1}[0-9]{2}",
				Prefix: FileNamePrefix,
				Suffix: "mp4",
				Ignore: true,
				Scan:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

package check

import (
	"ascii/types"
	"reflect"
	"testing"
)

func TestUsage(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want types.Data
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Usage(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpressions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Expressions(tt.args.s)
			if got != tt.want {
				t.Errorf("Expressions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Expressions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

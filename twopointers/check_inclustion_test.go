package twopointers

import (
	"reflect"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case -2",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
		{
			name: "case -1",
			args: args{
				s1: "abc",
				s2: "bbbca",
			},
			want: true,
		},
		{
			name: "case 0",
			args: args{
				s1: "hello",
				s2: "ooolleoooleh",
			},
			want: false,
		},
		{
			name: "case 1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				s1: "a",
				s2: "ab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeMap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				s: "abc",
			},
			want: map[string]int{
				"a": 1,
				"b": 1,
				"c": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeMap(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

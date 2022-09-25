package twopointers

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{
				s: []string{"h", "e", "l", "l", "o"},
			},
			want: []string{"o", "l", "l", "e", "h"},
		}, {
			name: "Case 2",
			args: args{
				s: []string{"H", "a", "n", "n", "a", "h"},
			},
			want: []string{"h", "a", "n", "n", "a", "H"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package twopointers

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

		{
			name: "Case 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{0, 0, 1},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{0, 1},
			},
			want: []int{1, 0},
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{1, 0, 1},
			},
			want: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

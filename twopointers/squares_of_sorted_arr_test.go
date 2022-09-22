package twopointers

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		// TODO: Add test cases.
		{
			name: "Case 2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

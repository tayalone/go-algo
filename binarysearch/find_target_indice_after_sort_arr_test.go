package binarysearch

import (
	"reflect"
	"testing"
)

func TestTargetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{1, 2, 5, 2, 3},
				target: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "Case 2",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 3,
			},
			want: []int{3},
		},
		{
			name: "Case 3",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 5,
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TargetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TargetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

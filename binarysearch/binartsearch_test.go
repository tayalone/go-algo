package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case: O Slice",
			args: args{
				nums:   []int{},
				target: 9,
			},
			want: -1,
		},
		{
			name: "Case: 1 Element & found",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "Case: 1 Element & not found",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "Case: 2 Elemnt & found at 1st ele",
			args: args{
				nums:   []int{1, 2},
				target: 1,
			},
			want: 0,
		},
		{
			name: "Case: 2 Elemnt & found at 2nd ele",
			args: args{
				nums:   []int{1, 2},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Case: 2 Elemnt & not found",
			args: args{
				nums:   []int{1, 2},
				target: 9,
			},
			want: -1,
		},
		{
			name: "Case: Found",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "Case: Not Found",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

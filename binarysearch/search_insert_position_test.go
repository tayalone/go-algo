package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1:",
			want: 2,
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
		}, {
			name: "Case 2:",
			want: 1,
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
		}, {
			name: "Case 3:",
			want: 4,
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
		}, {
			name: "Case 4:",
			want: 0,
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

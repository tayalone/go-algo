package binarysearch

import "testing"

func TestMySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Case 0",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "Case 1",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "Case 2",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "Case 3",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				x: 100,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySqrt(tt.args.x); got != tt.want {
				t.Errorf("MySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

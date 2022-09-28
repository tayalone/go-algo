package helper

import (
	"testing"

	"github.com/tayalone/go-algo/ds"
)

func TestCountTreeNodesVal(t *testing.T) {
	type args struct {
		t *ds.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				&ds.TreeNode{
					Val: 3,
					Left: &ds.TreeNode{
						Val: 4,
						Left: &ds.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &ds.TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &ds.TreeNode{
						Val:  5,
						Left: nil,
						Right: &ds.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 28,
		},
		{
			name: "case 2",
			args: args{
				&ds.TreeNode{
					Val: 2,
					Left: &ds.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTreeNodesVal(tt.args.t); got != tt.want {
				t.Errorf("CountTreeNodesVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

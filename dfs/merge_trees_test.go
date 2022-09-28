package dfs

import (
	"testing"

	"github.com/tayalone/go-algo/ds"
	"github.com/tayalone/go-algo/helper"
)

func TestMergeTrees(t *testing.T) {
	type args struct {
		root1 *ds.TreeNode
		root2 *ds.TreeNode
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
				root1: &ds.TreeNode{
					Val: 1,
					Left: &ds.TreeNode{
						Val: 3,
						Left: &ds.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &ds.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				root2: &ds.TreeNode{
					Val: 2,
					Left: &ds.TreeNode{
						Val:  1,
						Left: nil,
						Right: &ds.TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &ds.TreeNode{
						Val:  3,
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
				root1: &ds.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				root2: &ds.TreeNode{
					Val: 1,
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
			nt := MergeTrees(tt.args.root1, tt.args.root2)
			got := helper.CountTreeNodesVal(nt)

			if got != tt.want {
				t.Errorf("CountTreeNodesVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

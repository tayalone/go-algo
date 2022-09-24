package binarysearch

import "testing"

func TestEvaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				&TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				&TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateTree(tt.args.root); got != tt.want {
				t.Errorf("EvaluateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

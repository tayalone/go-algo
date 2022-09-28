package dfs

import (
	"github.com/tayalone/go-algo/ds"
)

/*MergeTrees return merge 2 TreeNode*/
func MergeTrees(root1 *ds.TreeNode, root2 *ds.TreeNode) *ds.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	return &ds.TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  MergeTrees(root1.Left, root2.Left),
		Right: MergeTrees(root1.Right, root2.Right),
	}
}

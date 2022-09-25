package binarysearch

/*TreeNode is node*/
type TreeNode struct {
	Val   int // 1 -> true, 0->false, 2->or, 3->and
	Left  *TreeNode
	Right *TreeNode
}

/*
EvaluateTree return bool

leetcode
Runtime: 7 ms, faster than 97.85% of Go online submissions for Evaluate Boolean Binary Tree.
Memory Usage: 6.4 MB, less than 87.10% of Go online submissions for Evaluate Boolean Binary Tree.
*/
func EvaluateTree(root *TreeNode) bool {
	if root == nil || root.Val == 0 {
		return false
	} else if root.Val == 1 {
		return true
	} else if root.Val == 2 {
		if (root.Left != nil && root.Left.Val == 1) || (root.Right != nil && root.Right.Val == 1) {
			return true
		}
		return EvaluateTree(root.Left) || EvaluateTree(root.Right)
	} else {
		if (root.Left != nil && root.Left.Val == 1) && (root.Right != nil && root.Right.Val == 1) {
			return true
		}
		return EvaluateTree(root.Left) && EvaluateTree(root.Right)
	}
}

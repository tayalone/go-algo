package helper

import (
	"github.com/tayalone/go-algo/ds"
)

/*CountTreeNodesVal return all value in truee node*/
func CountTreeNodesVal(t *ds.TreeNode) int {
	if t == nil {
		return 0
	}
	return t.Val + CountTreeNodesVal(t.Left) + CountTreeNodesVal(t.Right)
}

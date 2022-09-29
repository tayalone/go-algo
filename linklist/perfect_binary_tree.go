package linklist

// *ds.TreeNode
import (
	"github.com/tayalone/go-algo/ds"
)

/*Connect is return Perfect  */
func Connect(root *ds.PerfectNode) *ds.PerfectNode {
	// if root.l'
	if root == nil {
		return nil
	}

	cur, next := root, root.Left

	for cur != nil && next != nil {
		cur.Left.Next = cur.Right
		if cur.Next != nil {
			cur.Right.Next = cur.Next.Left
		}
		cur = cur.Next
		if cur == nil {
			cur = next
			next = cur.Left
		}
	}

	return root
}

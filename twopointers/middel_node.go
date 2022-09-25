package twopointers

/*MiddleNode is Reuturn  *ListNode */
func MiddleNode(head *ListNode) *ListNode {
	c := 1
	r := head

	for head.Next != nil {
		if c == 1 || (c > 2 && c%2 == 1) {
			r = r.Next
		}
		c++
		head = head.Next
	}
	return r
}

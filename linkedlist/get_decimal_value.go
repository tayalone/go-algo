package linkedlist

func getDecimalValue(head *ListNode) int {
	r := 0
	for head != nil {
		r = r + head.Val
		if head.Next != nil {
			r = r << 1
		}
		head = head.Next
	}
	return r
}

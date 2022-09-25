package twopointers

/*
RemoveNthFromEnd return list with remove n element
Do With Myself xoxo
LeetCode
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 31.49% of Go online submissions for Remove Nth Node From End of List.
*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	s := head
	f := head
	c := head
	counter := 0
	// sCounter := 0

	for c != nil {
		if c.Next == nil {
			if s == f {
				head = s.Next
			} else {
				s.Next = f.Next
			}
			break
		}
		c = c.Next
		if counter >= n-1 {
			f = f.Next
		}
		if counter >= n {
			s = s.Next
		}
		counter++
	}

	return head
}

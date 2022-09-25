package twopointers

/*RemoveNthFromEnd return list with remove n element  */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// cLink := head.Next
	// fLink := head.Next
	// lLink := head

	// counter := 2

	// for cLink != nil {
	// 	if cLink.Next == nil {
	// 		if counter == n {
	// 			head = head.Next
	// 			break
	// 		} else {
	// 			lLink.Next = fLink.Next
	// 			break
	// 		}
	// 	}
	// 	cLink = cLink.Next
	// 	counter++
	// }

	return head

}

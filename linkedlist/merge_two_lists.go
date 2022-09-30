package linkedlist

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

/*MergeTwoLists reuturn sorting *ListNode */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var r *ListNode = nil
	var c *ListNode = nil
	for list1 != nil || list2 != nil {
		var tmp *ListNode = nil
		bothExisting := list1 != nil && list2 != nil
		if list1 != nil && list2 == nil || bothExisting && list1.Val <= list2.Val {
			tmp = list1
			list1 = list1.Next
		} else if list1 == nil && list2 != nil || bothExisting && list1.Val > list2.Val {
			tmp = list2
			list2 = list2.Next
		}
		if r == nil {
			r = tmp
			c = tmp
		} else {
			c.Next = tmp
			c = tmp
		}
	}
	return r
}

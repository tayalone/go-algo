package twopointers

func rotate(nums []int, k int) []int {
	n := len(nums)
	realShift := k

	if k >= n {
		realShift = k % n
	}

	if realShift > 0 {
		return nums
	}

	return append(nums[n-realShift:], nums[:n-realShift]...)
}

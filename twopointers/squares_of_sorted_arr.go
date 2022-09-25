package twopointers

/*SortedSquares sort non-decreasing int slice */
func SortedSquares(nums []int) []int {
	lIndex := 0
	rIndex := len(nums) - 1
	r := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		lv := nums[lIndex] * nums[lIndex]
		rv := nums[rIndex] * nums[rIndex]
		if lv > rv {
			r[i] = lv
			lIndex++
		} else {
			r[i] = rv
			rIndex--
		}
	}
	return r
}

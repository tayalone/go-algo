package twopointers

/*MoveZeroes return []int which zero laydown @ end of slide*/
func MoveZeroes(nums []int) []int {
	ptr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[ptr]
			nums[ptr] = nums[i]
			nums[i] = tmp
			ptr++
		}
	}

	return nums
}

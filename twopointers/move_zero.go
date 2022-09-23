package twopointers

/*MoveZeroes return []int which zero laydown @ end of slide*/
func MoveZeroes(nums []int) []int {
	s := 0
	l := len(nums) - 1

	for s < l {
		if nums[s] == 0 {
			for i := s; i < l; i++ {
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
			l--
		}
		if nums[s] > 0 {
			s++
		}
	}

	return nums
}

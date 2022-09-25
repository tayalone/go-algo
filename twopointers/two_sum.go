package twopointers

/*TwoSum sort non-decreasing int slice */
func TwoSum(numbers []int, target int) []int {
	r := []int{}
	s := 0
	l := len(numbers) - 1

	for s < l {
		if numbers[s]+numbers[l] == target {
			return []int{s + 1, l + 1}
		} else if numbers[s]+numbers[l] < target {
			s++
		} else {
			l--
		}
	}

	return r
}

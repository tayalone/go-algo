package binarysearch

/*MissingNumber do find missingNumber in slice */
func MissingNumber(nums []int) int {
	result := len(nums)

	var intSet = make(map[int]bool)
	for i := range nums {
		intSet[i] = false
	}

	for _, value := range nums {
		_, ok := intSet[value]
		if ok {
			intSet[value] = true
		}
	}

	for key, value := range intSet {
		if !value {
			return key
		}
	}

	return result
}

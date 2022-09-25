package binarysearch

/*MissingNumber do find missingNumber in slice */
func MissingNumber(nums []int) int {
	sum := 0
	for i, v := range nums {
		sum = sum ^ i ^ v
	}
	return sum ^ len(nums)
}

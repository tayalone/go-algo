package binarysearch

/*
SearchInsert serch or insert
*/
func SearchInsert(nums []int, target int) int {
	sIndex := 0
	lIndex := len(nums) - 1

	for sIndex <= lIndex {
		pivot := (lIndex + sIndex) / 2
		if nums[pivot] == target {
			return pivot
		} else if target < nums[pivot] {
			lIndex = pivot - 1
		} else {
			sIndex = pivot + 1
		}
	}

	return sIndex
}

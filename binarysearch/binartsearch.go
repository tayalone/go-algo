package binarysearch

/*
BinarySearch : Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.

This function Performance
Runtime: 27 ms, faster than 98.81% of Go online submissions for Binary Search.
Memory Usage: 6.6 MB, less than 94.35% of Go online submissions for Binary Search.
21/09/2022
*/
func BinarySearch(nums []int, target int) int {
	sIndex := 0
	lIndex := len(nums) - 1
	index := -1
	for {
		cIndex := (sIndex + lIndex) / 2
		if sIndex > lIndex {
			break
		}
		if nums[cIndex] != target {
			if nums[cIndex] > target {
				lIndex = cIndex - 1
			} else {
				sIndex = cIndex + 1
			}
		} else {
			index = cIndex
			break
		}
	}

	return index
}

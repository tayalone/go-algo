package binarysearch

/*
BinarySearch : Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.

This function Performance
Runtime: 78 ms, faster than 13.89% of Go online submissions for Binary Search.
Memory Usage: 6.9 MB, less than 61.56% of Go online submissions for Binary Search.
21/09/2022
*/
func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums)
	sudoLen := end
	index := -1
	for {
		if sudoLen <= 0 {
			break
		} else if sudoLen == 1 {
			if nums[start] == target {
				index = start
			}
			break
		} else if sudoLen == 2 {
			if nums[start] == target {
				index = start
			} else if nums[start+1] == target {
				index = start + 1
			}
			break
		} else {
			isOdd := sudoLen%2 == 1
			middle1 := start + (sudoLen / 2)
			if isOdd {
				if nums[middle1] == target {
					index = middle1
					break
				} else if target < nums[middle1] {
					// / focus left side of slice
					sudoLen = middle1 - start
					end = middle1
				} else {
					// / focus right side of slice
					sudoLen = (end - middle1) + 1
					start = middle1
				}
			} else {
				middle2 := middle1 + 1
				if nums[middle1] == target {
					index = middle1
					break
				} else if nums[middle2] == target {
					index = middle2
					break
				} else if target < nums[middle1] {
					// / focus left side of slice
					sudoLen = middle1 - start
					end = middle1
				} else {
					// / focus right side of slice
					sudoLen = end - middle2
					start = middle2
				}
			}
		}
	}
	return index
}

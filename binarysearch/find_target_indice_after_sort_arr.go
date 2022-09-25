package binarysearch

import (
	"sort"
)

// /*
// TargetIndices return slice of target index
// */
// func TargetIndices(nums []int, target int) []int {
// 	s := sorts.BinaryInsertionSort(nums)

// 	f := func(arr []int, t int) []int {
// 		// enableGoRight := false
// 		r := []int{}

// 		for i := 0; i < len(arr); i++ {
// 			if arr[i] == target {
// 				r = append(r, i)
// 			} else if arr[i] > target {
// 				break
// 			}
// 		}

// 		return r
// 	}
// 	fmt.Println(s, target)
// 	r := f(s, target)
// 	fmt.Println(r)

// 	return f(nums, target)
// }

/*
TargetIndices return slice of target index
Write 1st time
*/
func TargetIndices(nums []int, target int) []int {
	res := []int{}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		pivot := (r + l) / 2
		if nums[pivot] < target {
			l = pivot + 1
		} else if nums[pivot] > target {
			r = pivot - 1
		} else {
			if nums[l] < target {
				l++
			}
			if nums[r] > target {
				r--
			}
			if nums[l] == target && nums[r] == target {
				res = append(res, l)
				l++
			}
		}
	}
	return res
}

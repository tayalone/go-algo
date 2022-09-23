package binarysearch

import (
	"fmt"

	"github.com/tayalone/go-algo/sorts"
)

/*
TargetIndices return slice of target index
Write 1st time
*/
func TargetIndices(nums []int, target int) []int {
	s := sorts.BinaryInsertionSort(nums)

	f := func(arr []int, t int) []int {
		// enableGoRight := false
		r := []int{}

		for i := 0; i < len(arr); i++ {
			if arr[i] == target {
				r = append(r, i)
			} else if arr[i] > target {
				break
			}
		}

		return r
	}
	fmt.Println(s, target)
	r := f(s, target)
	fmt.Println(r)

	return f(nums, target)
}

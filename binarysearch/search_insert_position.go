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

func sortedSquares(nums []int) []int {
	n := []int{}
	p := []int{}

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		isN := v < 0
		if isN {
			if len(n) == 0 {
				n = append(n, v)
			} else {
				n = append([]int{v}, n...)
			}
		} else {
			p = append(p, v)
		}

	}
	return append(p, n...)
}

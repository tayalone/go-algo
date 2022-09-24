package binarysearch

import "sort"

/*
Intersect2Arr return []int{}
leetCode (selfmade)
Runtime: 8 ms, faster than 30.35% of Go online submissions for Intersection of Two Arrays.
Memory Usage: 3 MB, less than 80.13% of Go online submissions for Intersection of Two Arrays.
*/
func Intersect2Arr(nums1 []int, nums2 []int) []int {
	r := []int{}
	m := make(map[int]bool)

	sort.Ints(nums1)
	sort.Ints(nums2)

	n1 := 0
	n2 := 0

	for n1 < len(nums1) && n2 < len(nums2) {
		repeat1 := m[nums1[n1]]
		repeat2 := m[nums2[n2]]

		if repeat1 == true && repeat2 == true {
			n1++
			n2++
			continue
		} else {
			if repeat1 == false && repeat2 == false && nums1[n1] == nums2[n2] {
				m[nums1[n1]] = true
				r = append(r, nums1[n1])
				n1++
				n2++
				continue
			} else {
				if nums1[n1] > nums2[n2] {
					n2++
					continue
				}
				n1++
			}
		}
	}

	return r
}

/*
HyperIntersect2Arr return []int{}
leetCode (cheat)
Runtime: 3 ms, faster than 84.06% of Go online submissions for Intersection of Two Arrays.
Memory Usage: 3.1 MB, less than 63.32% of Go online submissions for Intersection of Two Arrays.
*/
func HyperIntersect2Arr(nums1 []int, nums2 []int) []int {
	r := []int{}
	m := make(map[int]bool)

	if len(nums1) < len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}

	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		found := m[v]
		if found {
			r = append(r, v)
			delete(m, v)
		}
	}
	return r
}

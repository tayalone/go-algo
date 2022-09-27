package twopointers

import (
	"reflect"
)

func makeMap(s string) map[string]int {
	m := make(map[string]int)

	for i := 0; i < len(s); i++ {
		key := string(s[i])
		if value, found := m[key]; found {
			m[key] = value + 1
		} else {
			m[key] = 1
		}
	}
	return m
}

/*CheckInclusion is return true */
func CheckInclusion(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	m1 := makeMap(s1)

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		tmp := s2[i : i+len(s1)]
		m2 := makeMap(tmp)
		if reflect.DeepEqual(m1, m2) {
			return true
		}
	}

	return false
}

func cheateCheckInclusion(s1 string, s2 string) bool {

	// // create hashtable
	state1, state2 := make([]int, 26), make([]int, 26)
	matches := 0
	lS1 := len(s1)
	lS2 := len(s2)
	// //
	if lS1 > lS2 {
		return false
	}
	// // add counter in hash table
	for i := range s1 {
		state1[s1[i]-'a']++ // add array index of  s1[i]
		state2[s2[i]-'a']++ // add array index of  s2[i]
	}

	// // loop 26 round add matches++
	for i := range state1 {
		if state1[i] == state2[i] {
			matches++
		}
	}

	l, r := 0, lS1
	for r < lS2 {
		if matches == 26 {
			return true
		}
		rIndex := s2[r] - 'a'
		state2[rIndex]++
		if state2[rIndex]-1 == state1[rIndex] {
			matches--
		}
		if state2[rIndex] == state1[rIndex] {
			matches++
		}

		lIndex := s2[l] - 'a'

		state2[lIndex]--
		if state2[lIndex]+1 == state1[lIndex] {
			matches--
		}
		if state2[lIndex] == state1[lIndex] {
			matches++
		}
		l++
		r++
	}
	return matches == 26
}

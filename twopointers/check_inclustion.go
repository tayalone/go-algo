package twopointers

import "reflect"

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
	state1, state2 := make([]int, 26), make([]int, 26)
	matches := 0
	if len(s1) > len(s2) {
		return false
	}
	for i := range s1 {
		state1[s1[i]-'a']++
		state2[s2[i]-'a']++
	}
	for i := range state1 {
		if state1[i] == state2[i] {
			matches++
		}
	}
	l, r := 0, len(s1)
	for r < len(s2) {
		if matches == 26 {
			return true
		}
		state2[s2[r]-'a']++
		if state2[s2[r]-'a']-1 == state1[s2[r]-'a'] {
			matches--
		}
		if state2[s2[r]-'a'] == state1[s2[r]-'a'] {
			matches++
		}

		state2[s2[l]-'a']--
		if state2[s2[l]-'a']+1 == state1[s2[l]-'a'] {
			matches--
		}
		if state2[s2[l]-'a'] == state1[s2[l]-'a'] {
			matches++
		}
		l++
		r++
	}
	return matches == 26
}

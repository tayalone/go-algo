package twopointers

/*LengthOfLongestSubstring is return int*/
func LengthOfLongestSubstring(s string) int {
	rng := len(s)
	if rng < 2 {
		return rng
	}
	runes := []rune(s)
	ht := make(map[string]int)
	ht[string(runes[0])] = 0
	startIndex := 0

	tr := []rune{runes[0]}
	counter := 1
	for fp := 1; fp <= rng-1; fp++ {
		key := runes[fp]
		value, foundFp := ht[string(key)]
		if !foundFp || (foundFp && (value < startIndex)) {
			ht[string(key)] = fp
			tr = append(tr, key)

			if len(tr) > counter {
				counter = len(tr)
			}
		} else {
			if startIndex == value {
				/* เจอซ้ำตัวแรกของ Poiner */
				ht[string(key)] = fp
				startIndex = startIndex + 1
				tr = append(runes[startIndex:fp], key)
			} else if value == fp-1 {
				/* ตรงกับตัวสุดท้ย */
				ht[string(key)] = fp
				startIndex = fp
				tr = []rune{key}
			} else {
				ht[string(key)] = value + 1
				startIndex = startIndex + 1
				if startIndex+1 <= rng-1 {
					tr = append(runes[startIndex+1:fp], key)
				}
			}
		}
	}
	return counter
}

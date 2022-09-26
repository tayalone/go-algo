package twopointers

/*
LengthOfLongestSubstring is return int
Runtime: 19 ms, faster than 41.42% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 4.7 MB, less than 28.20% of Go online submissions for Longest Substring Without Repeating Characters.
*/
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
			if value == startIndex {
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
				// ตรงกลาง
				if fp >= rng-1 {
					// ขอบ
					continue
				}
				if runes[fp] == runes[fp+1] {
					// ถ้าจุดที่ชี้กับถัดไปซ้ำกัน
					ht[string(key)] = fp + 1
					startIndex = fp + 1
					fp++
					tr = []rune{runes[startIndex]}
					continue
				}
				startIndex = value + 1
				ht[string(key)] = fp
				tr = append(runes[startIndex:fp], key)
			}
		}
	}
	return counter
}

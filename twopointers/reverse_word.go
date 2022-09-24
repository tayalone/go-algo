package twopointers

import (
	"strings"
)

func reverseStr(s string) string {
	b := []byte(s)
	f := 0
	l := len(b) - 1

	for f <= l {
		tmp := b[f]
		b[f] = b[l]
		b[l] = tmp
		f++
		l--
	}
	return string(b)
}

/*ReverseWords do reserverse strings in words*/
func ReverseWords(s string) string {
	sSlice := strings.Split(s, " ")
	r := ""
	for i, v := range sSlice {
		tmp := reverseStr(v)
		r += tmp
		if i < len(sSlice)-1 {
			r += " "
		}
	}
	return r
}

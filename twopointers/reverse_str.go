package twopointers

/*ReverseString return  []byte*/
func ReverseString(s []string) []string {
	f := 0
	l := len(s) - 1

	for f <= l {
		tmp := s[f]
		s[f] = s[l]
		s[l] = tmp
		f++
		l--
	}

	return s
}

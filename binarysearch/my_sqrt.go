package binarysearch

/*MySqrt return int*/
func MySqrt(x int) int {
	result := 0
	l := 0
	r := x

	for l <= r {
		pivot := (r + l) / 2
		pow := pivot * pivot

		if pow <= x && (pivot+1)*(pivot+1) > x {
			result = pivot
			break
		} else if pow > x && (pivot-1)*(pivot-1) < x {
			result = pivot - 1
			break
		} else if pow < x {
			l = pivot + 1
		} else {
			r = pivot - 1
		}
	}

	return result
}

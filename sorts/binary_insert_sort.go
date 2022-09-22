package sorts

func binaryInsertionSortSearch(arr []int, item, start, end int) int {
	for start < end {
		mid := (start + end) / 2
		if arr[mid] <= item {
			start = mid + 1
			continue
		}
		end = mid
	}
	return start
}

// BinaryInsertionSort performs a binary insertion sort on an
func BinaryInsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currentElement := arr[i]
		pos := binaryInsertionSortSearch(arr, currentElement, 0, i)
		j := i
		for j > pos {
			arr[j] = arr[j-1]
			j--
		}
		arr[pos] = currentElement
	}
	return arr
}

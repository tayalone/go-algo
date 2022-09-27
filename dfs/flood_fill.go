package dfs

/*FloodFill return [][]int */
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	ogValue := image[sr][sc]
	if color != ogValue {
		trans(&image, sr, sc, ogValue, color)

	}
	return image
}

func trans(image *[][]int, sr int, sc int, ogValue int, newValue int) {
	if (*image)[sr][sc] == ogValue {
		(*image)[sr][sc] = newValue
		if sc+1 < len((*image)[0]) {
			trans(image, sr, sc+1, ogValue, newValue)
		}
		if sc-1 >= 0 {
			trans(image, sr, sc-1, ogValue, newValue)
		}
		if sr-1 >= 0 {
			trans(image, sr-1, sc, ogValue, newValue)
		}
		if sr+1 < len((*image)) {
			trans(image, sr+1, sc, ogValue, newValue)
		}
	}
}

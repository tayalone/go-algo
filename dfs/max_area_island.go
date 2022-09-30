package dfs

/*
MaxAreaOfIsland return area of island as int
leetCode
Runtime: 20 ms, faster than 52.69% of Go online submissions for Max Area of Island.
Memory Usage: 4.9 MB, less than 88.64% of Go online submissions for Max Area of Island.
*/
func MaxAreaOfIsland(grid [][]int) int {
	maxValue := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			// -1 is discovered area
			// 0 is water area
			// 1 is land area
			newMax := findLand(&grid, r, c)
			if newMax > maxValue {
				maxValue = newMax
			}
		}
	}
	return maxValue
}

func findLand(grid *[][]int, r int, c int) int {

	v := (*grid)[r][c]
	(*grid)[r][c] = -1
	if v == 0 || v == -1 {
		return 0
	}
	eastArea := 0
	if c+1 < len((*grid)[0]) {
		eastArea = findLand(grid, r, c+1)
	}
	southArea := 0
	if r+1 < len((*grid)) {
		southArea = findLand(grid, r+1, c)
	}
	westArea := 0
	if c-1 >= 0 {
		westArea = findLand(grid, r, c-1)
	}
	northArea := 0
	if r-1 >= 0 {
		northArea = findLand(grid, r-1, c)
	}
	return 1 + eastArea + southArea + westArea + northArea
}

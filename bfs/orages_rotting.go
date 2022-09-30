package bfs

/*OrangesRotting return minute */
func OrangesRotting(grid [][]int) int {
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	r, c := len(grid), len(grid[0])
	var rottenQ [][2]int

	fresh := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				rottenQ = append(rottenQ, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(rottenQ) > 0 {
		for _, o := range rottenQ {
			// // วนลูป 4 directon
			for _, dir := range directions {
				y, x := o[0]+dir[0], o[1]+dir[1]
				if y >= 0 && y < r && x >= 0 && x < c && grid[y][x] == 1 {
					grid[y][x] = 2
					rottenQ = append(rottenQ, [2]int{y, x})
					fresh--
				}
			}
			rottenQ = rottenQ[1:]

		}
		minutes++
		if fresh == 0 {
			return minutes
		}
	}

	return -1
}

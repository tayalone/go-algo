package dynamicprograming

/*UpdateMatrix return nearest zero each index*/
func UpdateMatrix(mat [][]int) [][]int {
	// // ---------------- horizontal -----------------------------
	// if len(mat) <= 2 {
	// 	return mat
	// }
	// center := len(mat) / 2
	// // leftPtr := 1
	// rightPtr := len(mat) - 2
	// for leftPtr := 1; leftPtr <= center; leftPtr++ {
	// 	// // --------------------------------------
	// 	if leftPtr == rightPtr {
	// 		l := mat[leftPtr-1]
	// 		r := mat[leftPtr+1]
	// 		if r < l {
	// 			mat[leftPtr] = mat[leftPtr] + r
	// 		} else {
	// 			mat[leftPtr] = mat[leftPtr] + l
	// 		}
	// 		rightPtr--
	// 	} else {
	// 		if mat[leftPtr-1] > 0 {
	// 			mat[leftPtr] = mat[leftPtr] + mat[leftPtr-1]
	// 		}
	// 		if mat[rightPtr+1] > 0 {
	// 			mat[rightPtr] = mat[rightPtr] + mat[rightPtr+1]
	// 		}
	// 	}

	// 	// // ---------------------------------------

	// 	if rightPtr-1 == leftPtr {
	// 		break
	// 	}
	// 	rightPtr--
	// }
	// // --------------------------------------------------------------
	if len(mat) <= 2 {
		return mat
	}
	vCen := len(mat) / 2
	btmPtr := len(mat) - 1

	for topPtr := 0; topPtr <= vCen; topPtr++ {

		if topPtr == btmPtr {
			t := mat[topPtr-1][0]
			b := mat[topPtr+1][0]
			if t < b {
				mat[topPtr][0] = mat[topPtr][0] + mat[topPtr-1][0]
			} else {
				mat[topPtr][0] = mat[topPtr][0] + mat[topPtr+1][0]
			}
		} else {
			if topPtr > 0 && mat[topPtr-1][0] > 0 {
				mat[topPtr][0] = mat[topPtr][0] + mat[topPtr-1][0]
			}
			if btmPtr < len(mat)-1 && mat[btmPtr+1][0] > 0 {
				mat[btmPtr][0] = mat[btmPtr][0] + mat[btmPtr+1][0]
			}
		}

		if btmPtr-1 == topPtr {
			break
		}
		btmPtr--
	}

	return mat
}

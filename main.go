package main

import (
	"fmt"

	"github.com/tayalone/go-algo/binarysearch"
)

// nums: []int{1, 2, 3, 4, 5, 6, 7},
// 				k:    3,

func main() {
	binarysearch.TargetIndices(
		[]int{1, 2, 5, 2, 3},
		2,
	)
	fmt.Println("Hello Algo 101")
}

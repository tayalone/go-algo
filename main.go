package main

import (
	"fmt"

	"github.com/tayalone/go-algo/binarysearch"
	"github.com/tayalone/go-algo/binarysearch/mocks"
)

func main() {
	// fmt.Println(binarysearch.BinarySearch([]int{-1, 0, 5}, 5))
	myIbv := mocks.NewIbv(1)
	fbv := binarysearch.NewFBV(myIbv)

	fmt.Println(myIbv.IsBadVersion(3))

	fmt.Println(fbv.FirstBadVersion(3))

	fmt.Println("Hello Algo 101")
}

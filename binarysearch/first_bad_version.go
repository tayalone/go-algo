package binarysearch

import (
	"github.com/tayalone/go-algo/binarysearch/mocks"
)

/*FBVStruct is struct*/
type FBVStruct struct {
	ibv mocks.IBV
}

/*FBV is interface*/
type FBV interface {
	FirstBadVersion(n int) int
}

/*NewFBV is Instance*/
func NewFBV(ibv mocks.IBV) FBVStruct {
	return FBVStruct{ibv}
}

var defaltBadVersion = 1

/*FirstBadVersion return FirstBadVersion*/
func (f FBVStruct) FirstBadVersion(n int) int {
	m := 1
	for m < n {
		pivot := (n + m) / 2
		if f.ibv.IsBadVersion(pivot) {
			n = pivot
			continue
		}
		m = pivot + 1
	}
	return n
}

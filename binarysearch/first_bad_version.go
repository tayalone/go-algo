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
	find := func(start, end int, findIsBad bool) int {
		ts := start
		te := end
		for {
			tv := (ts + te) / 2
			isBad := f.ibv.IsBadVersion(tv)
			// /*
			// 		ถ้า until เป็น true
			// 		 - แล้ว r เป็น false ขยับ ts
			// 		 - ถ้าเป็น trur return tv
			// 	    ถ้า until เป็น false
			// 		 - แล้ว r เป็น true ขยับ te
			// 		 - ถ้าเป็น false return tv
			// */
			// if until {}

			/* ต้่องการหา fail version  */

			if findIsBad == isBad {
				return tv
			} else if findIsBad {
				ts = tv + 1
			} else {
				te = te - 1
			}
		}
	}
	findIsBad := true
	gv := 1
	bv := n

	if f.ibv.IsBadVersion(1) {
		return 1
	}

	for {
		r := find(gv, bv, findIsBad)
		if findIsBad {
			bv = r
		} else {
			gv = r
		}
		if (bv - gv) == 1 {
			return bv
		}
		findIsBad = !findIsBad
	}
}

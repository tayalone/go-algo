package binarysearch

/*FBVStruct is struct*/
type FBVStruct struct{}

/*FBV is interface*/
type FBV interface {
	firstBadVersion(n int) int
}

/*MyFbv is Instance*/
var MyFbv = FBVStruct{}

func (f FBVStruct) firstBadVersion(n int) int {
	return n
}

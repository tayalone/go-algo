package mocks

/*IBVStruct is struct*/
type IBVStruct struct {
	version int
}

/*IBV is interface*/
type IBV interface {
	IsBadVersion(version int) bool
}

/*NewIbv is Instance*/
func NewIbv(version int) IBVStruct {
	return IBVStruct{
		version: version,
	}
}

/*IsBadVersion do return 1st bad version*/
func (ibv IBVStruct) IsBadVersion(version int) bool {
	return ibv.version <= version
}

package uuid

import "fmt"

var mockedUUIDCounter uint32 = 0
var enableMockedUUID = false

func EnableMockUUID(start_count uint32) {
	mockedUUIDCounter = start_count
	enableMockedUUID = true
}
func DisableMockUUID() {
	enableMockedUUID = false
}
func NewMockUUID(count uint32) UUID {
	return UUID(fmt.Sprintf("ffffffff-ffff-ffff-ffff-%012d", count))
}

func generateMockedUUID() UUID {
	u := NewMockUUID(mockedUUIDCounter)
	mockedUUIDCounter += 1
	return u
}

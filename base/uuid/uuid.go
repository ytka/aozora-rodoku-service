package uuid

import google_uuid "github.com/google/uuid"

type UUID string

func generateUUID() UUID {
	return UUID(google_uuid.NewString())
}

func New() UUID {
	if enableMockedUUID {
		return generateMockedUUID()
	} else {
		return generateUUID()
	}
}

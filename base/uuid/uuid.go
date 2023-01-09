package uuid

import google_uuid "github.com/google/uuid"

type UUID string

func New() UUID {
	return UUID(google_uuid.NewString())
}

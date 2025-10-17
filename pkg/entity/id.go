package id

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId() ID {
	id := uuid.New()

	return ID(id)
}

func ParseId(id string) (ID, error) {
	parsedId, err := uuid.Parse(id)

	return ID(parsedId), err
}

package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

var (
	ErrInvalidID = errors.New("invalid id")
)

type ID string

func NewID(id string) (ID, error) {
	if len(id) < 6 {
		return "", fmt.Errorf("%w: %s", ErrInvalidID, id)
	}

	return ID(id), nil
}

var GenerateUUID = generateUUID

func generateUUID() ID {
	return ID(uuid.New().String())
}

func PatchGenerateUUIDWith(patch func() ID) {
	GenerateUUID = patch
}

func RestoreGenerateUUID() {
	GenerateUUID = generateUUID
}

func GenerateSlug(str string) (ID, error) {
	s := slug.Make(str)
	return NewID(s)
}

func (id ID) String() string {
	return string(id)
}

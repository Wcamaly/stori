package models

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidRequiredString = errors.New("invalid required string")
)

type RequiredString string

func NewRequiredString(str string) (RequiredString, error) {
	if len(strings.TrimSpace(str)) < 1 {
		return "", fmt.Errorf("%w: %s", ErrInvalidRequiredString, str)
	}

	return RequiredString(str), nil
}

func (id RequiredString) Value() string {
	return string(id)
}

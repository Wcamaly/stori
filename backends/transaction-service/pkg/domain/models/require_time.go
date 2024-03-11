package models

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidRequiredTime = errors.New("invalid required time")
)

type RequiredTime time.Time

func NewRequiredTime(str string) (*RequiredTime, error) {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidRequiredTime, str)
	}

	return (*RequiredTime)(&t), nil
}

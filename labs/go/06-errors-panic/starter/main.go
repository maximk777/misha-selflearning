package starter

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("user not found")

func FindUser(id string) (string, error) {
	return "", fmt.Errorf("find user %q: %w", id, ErrNotFound)
}

func MustPositive(value int) int {
	if value <= 0 {
		panic("value must be positive")
	}
	return value
}

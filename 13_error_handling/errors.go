package main

import (
	"errors"
	"fmt"
)

var ErrNegativeNumber = errors.New("negative number")

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func safeSqrt(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("safeSqrt(%d): %w", n, ErrNegativeNumber)
	}

	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return i, nil
		}
	}

	return 0, &ValidationError{
		Field: "n",
		Msg:   "not a perfect square",
	}
}

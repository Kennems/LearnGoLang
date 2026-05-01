package main

import (
	"context"
	"errors"
	"time"
)

func waitOrCancel(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runWithTimeout(parent context.Context, timeout time.Duration, work time.Duration) error {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	return waitOrCancel(ctx, work)
}

func isDeadlineExceeded(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}

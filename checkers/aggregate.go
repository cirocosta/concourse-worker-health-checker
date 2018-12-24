package checkers

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Aggregate struct {
	Checkers []Checker
}

func (h *Aggregate) Check(ctx context.Context) (err error) {
	var group *errgroup.Group

	group, ctx = errgroup.WithContext(ctx)
	for _, checker := range h.Checkers {
		checker := checker // goroutine closure

		group.Go(func() error {
			return checker.Check(ctx)
		})
	}

	err = group.Wait()
	return
}

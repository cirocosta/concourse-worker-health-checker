package healthcheck_test

import (
	"context"
	"fmt"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type noopChecker struct {
	WaitCancellation bool
	Fail             bool

	timesCalled  int
	wasCancelled bool
}

func (c *noopChecker) Check(ctx context.Context) (err error) {
	if c.WaitCancellation {
		<-ctx.Done()
		c.wasCancelled = true
	}

	c.timesCalled++

	if c.Fail {
		err = fmt.Errorf("aa")
		return
	}

	return
}

func (c *noopChecker) TimesCalled() int {
	return c.timesCalled
}

func (c *noopChecker) WasCancelled() bool {
	return c.wasCancelled
}

var _ = Describe("aggregate", func() {
	var (
		aggregate *healthcheck.Aggregate
		err       error
	)

	JustBeforeEach(func() {
		err = aggregate.Check(context.Background())
	})

	Context("having some checkers", func() {
		var registeredCheckers []healthcheck.Checker

		Context("with none of them erroring", func() {
			BeforeEach(func() {
				registeredCheckers = []healthcheck.Checker{
					&noopChecker{},
					&noopChecker{},
					&noopChecker{},
				}

				aggregate = &healthcheck.Aggregate{
					Checkers: registeredCheckers,
				}
			})

			It("eventually gets them all called", func() {
				Eventually(func() int {
					checked := 0

					for _, checker := range registeredCheckers {
						checked += checker.(*noopChecker).TimesCalled()
					}

					return checked
				}).Should(Equal(len(registeredCheckers)))
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with at least one failing", func() {
			BeforeEach(func() {
				registeredCheckers = []healthcheck.Checker{
					&noopChecker{Fail: true},
					&noopChecker{WaitCancellation: true},
					&noopChecker{WaitCancellation: true},
				}

				aggregate = &healthcheck.Aggregate{
					Checkers: registeredCheckers,
				}
			})

			It("fails the check", func() {
				Expect(err).To(HaveOccurred())
			})

			It("doesn't call them all", func() {
				Expect(registeredCheckers[1].(*noopChecker).WasCancelled()).To(BeTrue())
				Expect(registeredCheckers[2].(*noopChecker).WasCancelled()).To(BeTrue())
			})
		})
	})
})

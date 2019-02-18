package healthcheck_test

import (
	"context"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"

	fakes "github.com/cirocosta/concourse-worker-health-checker/healthcheck/healthcheckfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Worker", func() {
	var (
		containerProvider *fakes.FakeContainerProvider
		volumeProvider    *fakes.FakeVolumeProvider
		worker            *healthcheck.Worker
		err               error
	)

	BeforeEach(func() {
		containerProvider = &fakes.FakeContainerProvider{}
		volumeProvider = &fakes.FakeVolumeProvider{}

		volumeProvider.CreateReturns(&healthcheck.Volume{
			Handle: "handle",
			Path:   "/rootfs",
		}, nil)

		worker = &healthcheck.Worker{
			ContainerProvider: containerProvider,
			VolumeProvider:    volumeProvider,
		}
	})

	Context("Check", func() {
		JustBeforeEach(func() {
			err = worker.Check(context.TODO())
		})

		Context("having volume and container creation working", func() {
			It("doesn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("calls volume creation", func() {
				Expect(volumeProvider.CreateCallCount()).To(Equal(1))
			})

			It("calls volume deletion", func() {
				Expect(volumeProvider.DestroyCallCount()).To(Equal(1))
			})

			It("calls container creation", func() {
				Expect(containerProvider.CreateCallCount()).To(Equal(1))
			})

			It("calls container deletion", func() {
				Expect(containerProvider.DestroyCallCount()).To(Equal(1))
			})
		})

		// TODO explore the possible errors
	})
})

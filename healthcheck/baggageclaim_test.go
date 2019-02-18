package healthcheck_test

import (
	"context"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("baggageclaim", func() {
	var (
		bcServer *ghttp.Server
		bc       *healthcheck.Baggageclaim
		err      error
		volume   *healthcheck.Volume
	)

	BeforeEach(func() {
		bcServer = ghttp.NewServer()
		bc = &healthcheck.Baggageclaim{
			Url: "http://" + bcServer.Addr(),
		}
	})

	Context("Create", func() {
		var statusCode = 200

		JustBeforeEach(func() {
			volume, err = bc.Create(context.TODO(), "handle")
		})

		BeforeEach(func() {
			expectedVol := healthcheck.Volume{Handle: "handle", Path: "/rootfs"}

			bcServer.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/volumes"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &expectedVol),
			))
		})

		It("issues volume creation request", func() {
			Expect(bcServer.ReceivedRequests()).To(HaveLen(1))
		})

		Context("having positive response", func() {
			It("doesn't fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns a proper volume", func() {
				Expect(volume.Handle).To(Equal("handle"))
				Expect(volume.Path).To(Equal("/rootfs"))
			})
		})

		Context("having negative response", func() {
			BeforeEach(func() {
				statusCode = 500
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("Destroy", func() {
		var statusCode = 200

		JustBeforeEach(func() {
			err = bc.Destroy(context.TODO(), "handle")
		})

		BeforeEach(func() {
			bcServer.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("DELETE", MatchRegexp(`/volumes/[a-z0-9-]+`)),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, nil),
			))
		})

		It("issues volume deletion request", func() {
			Expect(bcServer.ReceivedRequests()).To(HaveLen(1))
		})
	})
})

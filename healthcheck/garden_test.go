package healthcheck_test

import (
	"context"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("garden", func() {
	var (
		gServer *ghttp.Server
		g       *healthcheck.Garden
		err     error
	)

	BeforeEach(func() {
		gServer = ghttp.NewServer()
		g = &healthcheck.Garden{
			Url: "http://" + gServer.Addr(),
		}
	})

	Context("Create", func() {
		var statusCode = 200

		JustBeforeEach(func() {
			err = g.Create(context.TODO(), "handle", "/rootfs")
		})

		BeforeEach(func() {
			gServer.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/containers"),
				ghttp.VerifyJSON(`{"handle":"handle","rootfs":"raw:///rootfs"}`),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, nil),
			))
		})

		It("issues container creation request", func() {
			Expect(gServer.ReceivedRequests()).To(HaveLen(1))
		})

		Context("having positive response", func() {
			It("doesn't fail", func() {
				Expect(err).ToNot(HaveOccurred())
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
			err = g.Destroy(context.TODO(), "handle")
		})

		BeforeEach(func() {
			gServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", MatchRegexp(`/containers/[a-z0-9-]+`)),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, nil),
				))
		})

		It("issues volume deletion request", func() {
			Expect(gServer.ReceivedRequests()).To(HaveLen(1))
		})

		Context("having positive response", func() {
			It("doesn't fail", func() {
				Expect(err).ToNot(HaveOccurred())
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
})

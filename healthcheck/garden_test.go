package healthcheck_test

import (
	"context"

	"github.com/concourse/concourse/worker/healthcheck"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("garden", func() {
	var (
		gServer *ghttp.Server
		err     error
	)

	BeforeEach(func() {
		gServer = ghttp.NewServer()
	})

	JustBeforeEach(func() {
		err = (&healthcheck.Garden{Url: "http://" + gServer.Addr()}).
			Check(context.Background())
	})

	Context("trying to create a volume", func() {
		Context("with the creation failing", func() {
			BeforeEach(func() {
				gServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/containers"),
						ghttp.RespondWith(500, "ok")))
			})

			It("issues creation request", func() {
				Expect(gServer.ReceivedRequests()).To(HaveLen(1))
			})

			It("returns error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("having the creation working", func() {
			BeforeEach(func() {
				gServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/containers"),
						ghttp.RespondWith(200, "ok")))

			})

			Context("having the deletion succeeding", func() {
				BeforeEach(func() {
					gServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("DELETE", MatchRegexp(`/containers/[a-z0-9-]+`)),
							ghttp.RespondWith(200, "ok")))
				})

				It("issues both requests", func() {
					Expect(gServer.ReceivedRequests()).To(HaveLen(2))
				})

				It("succeeds", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("having the deletion failing", func() {
				BeforeEach(func() {
					gServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("DELETE", MatchRegexp(`/containers/[a-z0-9-]+`)),
							ghttp.RespondWith(500, "ok")))
				})

				It("issues both requests", func() {
					Expect(gServer.ReceivedRequests()).To(HaveLen(2))
				})

				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})

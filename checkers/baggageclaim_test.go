package checkers_test

import (
	"context"

	"github.com/cirocosta/concourse-worker-health-checker/checkers"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("baggageclaim", func() {
	var (
		bcServer *ghttp.Server
		err      error
	)

	BeforeEach(func() {
		bcServer = ghttp.NewServer()
	})

	JustBeforeEach(func() {
		err = (&checkers.Baggageclaim{Address: "http://" + bcServer.Addr()}).
			Check(context.Background())
	})

	Context("trying to create a volume", func() {
		Context("with the creation failing", func() {
			BeforeEach(func() {
				bcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/volumes"),
						ghttp.RespondWith(500, "ok")))
			})

			It("issues creation request", func() {
				Expect(bcServer.ReceivedRequests()).To(HaveLen(1))
			})

			It("returns error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("having the creation working", func() {
			BeforeEach(func() {
				bcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/volumes"),
						ghttp.RespondWith(200, "ok")))

			})

			Context("having the deletion succeeding", func() {
				BeforeEach(func() {
					bcServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("DELETE", MatchRegexp(`/volumes/[a-z0-9-]+`)),
							ghttp.RespondWith(200, "ok")))
				})

				It("issues both requests", func() {
					Expect(bcServer.ReceivedRequests()).To(HaveLen(2))
				})

				It("succeeds", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("having the deletion failing", func() {
				BeforeEach(func() {
					bcServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest("DELETE", MatchRegexp(`/volumes/[a-z0-9-]+`)),
							ghttp.RespondWith(500, "ok")))
				})

				It("issues both requests", func() {
					Expect(bcServer.ReceivedRequests()).To(HaveLen(2))
				})

				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})

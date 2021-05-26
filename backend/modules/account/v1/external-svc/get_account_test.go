package external_svc_test

import (
	pb "alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Describe:GetAccount", func() {
	var (
		accountServer pb.AccountsServer
		ctx           context.Context
		account       *pb.Account
	)
	BeforeEach(func() {
		accountServer = AccountServerTest
		ctx = CtxTest
		account = Account
	})

	Describe("Get an account", func() {
		By("internal or external call")
		Context("Get an error when ctx is wrong", func() {
			It("Should get an error of permission denied", func() {

			})
		})

		Context("Get an error when id is wrong", func() {
			It("should return not found error", func() {
				_, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: "2e8a45a2-d29b-4764-b548-459db10e6a46"})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})

		Context("Get a record when id is provided", func() {
			It("should return requested field in the object", func() {
				res, err := accountServer.GetAccount(ctx, &pb.GetAccountRequest{Id: account.GetId()})
				Expect(err).Should(BeNil())
				Expect(res.Id).ShouldNot(BeEmpty())
				Expect(res.UserId).ShouldNot(BeEmpty())
			})
		})
	})
})

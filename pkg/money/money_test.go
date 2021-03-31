package money_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/simeonkorchev/rentim/pkg/money"
)

var _ = Describe("Money", func() {
	var (
		m money.Money
		err error
	)
	Describe("Creating EUR", func() {
		BeforeEach(func() {
			m, err = money.Make(1, money.EUR)
		})

		Context("1 EUR", func() {
			It("amount should be 1", func() {
				Expect(m.Amount()).To(Equal(1.0))
			})

			It("amount should be greater than zero", func() {
				Expect(m.LessThanOrEqualsZero()).To(BeFalse())
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	When("Currency is not expected", func() {
		BeforeEach(func() {
			m, err = money.Make(1, money.Currency{})
		})

		It("should return error", func() {
			Expect(err).To(MatchError("unknown currency ''"))
		})
	})

})

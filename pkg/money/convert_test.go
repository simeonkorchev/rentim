package money_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/simeonkorchev/rentim/pkg/money"
)

var _ = Describe("Convert", func() {
	var (
		m money.Money
		currency money.Currency
		amount float64
		err error
	)

	Describe("EUR to BGN", func() {
		JustBeforeEach(func() {
			m, _ = money.Make(amount, currency)
			m, err = money.ToBGN(m)
		})

		Context("1 EUR to BGN", func() {
			BeforeEach(func() {
				amount = 1
				currency = money.EUR
			})

			It("should be  1.9558 BGN", func() {
				Expect(m.Amount()).To(Equal(1.9558))
			})

			It("should not error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("2 EUR to BGN", func() {
			BeforeEach(func() {
				amount = 2
				currency = money.EUR
			})

			It("should equal to 2 * 1.9558 BGN", func() {
				Expect(m.Amount()).To(Equal(2 * 1.9558))
			})

			It("should not return error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("BGN to BGN", func() {
		JustBeforeEach(func() {
			m, _ = money.Make(amount, currency)
			m, err = money.ToBGN(m)
		})

		Context("1 BGN to BGN", func() {
			BeforeEach(func() {
				amount = 1
				currency = money.BGN
			})

			It("should return error", func() {
				Expect(err).To(MatchError("expected 'EUR' got 'BGN'"))
			})
		})
	})

	Describe("BGN to EUR", func() {
		JustBeforeEach(func() {
			m, _ = money.Make(amount, currency)
			m, err = money.ToEUR(m)
		})

		Context("1 BGN to EUR", func() {
			BeforeEach(func() {
				amount = 1
				currency = money.BGN
			})

			It("should be 0.511292", func() {
				Expect(m.Amount()).To(Equal(0.511292))
			})
			It("should not error", func() {
				Expect(err).To(BeNil())
			})
		})

	})

	Describe("EUR to EUR", func() {
		JustBeforeEach(func() {
			m, _ = money.Make(amount, currency)
			m, err = money.ToEUR(m)
		})

		Context("1 EUR to EUR", func() {
			BeforeEach(func() {
				amount = 1
				currency = money.EUR
			})

			It("should return error", func() {
				Expect(err).To(MatchError("expected 'BGN' got 'EUR'"))
			})
		})

	})
})

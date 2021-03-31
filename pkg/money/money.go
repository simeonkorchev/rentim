package money

import "fmt"

var (
	EUR = Currency{code: "EUR"}
	BGN = Currency{code: "BGN"}
)

type Currency struct {
	code string
}

type Money struct {
	amount  float64
	currency Currency
}

func (m Money) LessThanOrEqualsZero() bool {
	return m.amount <= 0
}

func (m Money) Amount() float64 {
	return m.amount
}

func Make(amount float64, currency Currency) (Money, error) {
	if currency != EUR && currency != BGN {
		return Money{}, fmt.Errorf("unknown currency '%v'", currency.code)
	}
	return Money{
		amount:   amount,
		currency: currency,
	}, nil
}
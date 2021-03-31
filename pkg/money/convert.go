package money

import "fmt"

func ToEUR(bgn Money) (Money, error) {
	if bgn.currency != BGN {
		return Money{}, fmt.Errorf("expected '%v' got '%v'", BGN.code, bgn.currency.code)
	}
	return Money{currency: BGN, amount: 0.511292 * bgn.amount}, nil
}

func ToBGN(euro Money) (Money, error) {
	if euro.currency != EUR {
		return Money{}, fmt.Errorf("expected '%v' got '%v'", EUR.code, euro.currency.code)
	}
	return Money{currency: BGN, amount: 1.9558 * euro.amount}, nil
}

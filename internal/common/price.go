package common

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Price struct {
	Value    decimal.Decimal
	Currency string
}

func NewPrice(value decimal.Decimal, currency string) *Price {
	return &Price{Value: value, Currency: currency}
}

func (price *Price) String() string {
	return fmt.Sprintf("%s%d", price.Currency, price.Value.CoefficientInt64())
}

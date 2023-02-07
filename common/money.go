package common

import "fmt"

const (
	DefaultCurrencyUnit                    CurrencyUnit = INR
	DefaultCurrencySubUnitConversionFactor int64        = 100
)

type Currency struct {
	valueInSubUnit          int64
	unit                    CurrencyUnit
	subUnitConversionFactor int64
}

type CurrencyUnit string

const (
	INR = "₹"
)

func NewCurrency(valueInSubUnit int64, unit CurrencyUnit, subUnitConversionFactor int64) (*Currency, error) {
	if subUnitConversionFactor == 0 {
		return nil, fmt.Errorf("error: subUnitConversionFactor cannot be 0")
	}
	return &Currency{valueInSubUnit, unit, subUnitConversionFactor}, nil
}

func (m *Currency) DisplayValue() string {
	return fmt.Sprintf("%s %.2f", m.unit, float64(m.valueInSubUnit)/float64(m.subUnitConversionFactor))
}

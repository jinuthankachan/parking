package common

import (
	"fmt"
)

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

func DefaultCurrency(amountInSubUnit int64) *Currency {
	return &Currency{
		valueInSubUnit:          amountInSubUnit,
		unit:                    DefaultCurrencyUnit,
		subUnitConversionFactor: DefaultCurrencySubUnitConversionFactor,
	}
}

func (m *Currency) Add(n *Currency) error {
	if m.unit != n.unit {
		return fmt.Errorf("invalid operation: %s(%d) + %s(%d) (mismatched units)",
			string(m.unit), m.valueInSubUnit,
			string(n.unit), n.valueInSubUnit)
	}
	m.valueInSubUnit += n.valueInSubUnit
	return nil
}

func (m *Currency) ValueInSubUnit() int64 {
	return m.valueInSubUnit
}
func (m *Currency) Unit() CurrencyUnit {
	return m.unit
}
func (m *Currency) SubUnitConversionFactor() int64 {
	return m.subUnitConversionFactor
}

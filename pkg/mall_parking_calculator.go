package pkg

import (
	"math"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type mallFeeCalculator struct {
	models.ParkingLotFees
}

func (calc *mallFeeCalculator) Calculate(
	vehicleType common.VehicleType,
	totalTimeParked time.Duration,
) (*common.Currency, error) {
	rate, err := calc.GetSlabForTime(totalTimeParked, common.Mall, vehicleType)
	if err != nil {
		return nil, err
	}
	chargeableHours := int64(math.Ceil(totalTimeParked.Hours()))
	return common.NewCurrency(
		rate.Charge.ValueInSubUnit()*chargeableHours,
		rate.Charge.Unit(),
		rate.Charge.SubUnitConversionFactor(),
	)
}

func MallFeeCalculator() *mallFeeCalculator {
	fees := models.NewParkingLotFees()
	return &mallFeeCalculator{
		*fees,
	}
}

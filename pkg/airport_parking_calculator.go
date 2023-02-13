package pkg

import (
	"math"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type airportFeeCalculator struct {
	models.ParkingLotFees
}

func (calc *airportFeeCalculator) Calculate(
	vehicleType common.VehicleType,
	totalTimeParked time.Duration,
) (*common.Currency, error) {
	rate, err := calc.GetSlabForTime(totalTimeParked, common.Airport, vehicleType)
	if err != nil {
		return nil, err
	}

	if rate.SlabType == common.PerDaySlab {
		chargeableDays := int64(math.Ceil(totalTimeParked.Hours() / 24))
		return common.NewCurrency(
			rate.Charge.ValueInSubUnit()*chargeableDays,
			rate.Charge.Unit(),
			rate.Charge.SubUnitConversionFactor(),
		)
	}
	return rate.Charge, nil
}

func AirportFeeCalculator() *airportFeeCalculator {
	fees := models.NewParkingLotFees()
	return &airportFeeCalculator{
		*fees,
	}
}

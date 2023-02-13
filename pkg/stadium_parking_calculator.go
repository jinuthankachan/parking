package pkg

import (
	"math"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type stadiumFeeCalculator struct {
	models.ParkingLotFees
}

func (calc *stadiumFeeCalculator) Calculate(
	vehicleType common.VehicleType,
	totalTimeParked time.Duration,
) (*common.Currency, error) {
	rates, err := calc.GetAllSlabsUnderTime(totalTimeParked, common.Stadium, vehicleType)
	if err != nil {
		return nil, err
	}

	var fees *common.Currency
	for _, rate := range rates {
		if fees == nil {
			fees, err = common.NewCurrency(0, rate.Charge.Unit(), rate.Charge.SubUnitConversionFactor())
			if err != nil {
				return nil, err
			}
		}

		if rate.SlabType == common.FlatSlab {
			fees.Add(rate.Charge)
		} else if rate.SlabType == common.PerHourSlab {
			var chargeableDurationInSlab time.Duration
			if totalTimeParked >= rate.UpperLimit {
				chargeableDurationInSlab = rate.UpperLimit - rate.LowerLimit
			} else {
				chargeableDurationInSlab = totalTimeParked - rate.LowerLimit
			}
			chargeableHours := int64(math.Ceil(math.Max(0.1, chargeableDurationInSlab.Hours())))
			chargesForSlab, err := common.NewCurrency(
				(chargeableHours)*rate.Charge.ValueInSubUnit(),
				rate.Charge.Unit(),
				rate.Charge.SubUnitConversionFactor(),
			)
			if err != nil {
				return nil, err
			}
			fees.Add(chargesForSlab)
		}
	}

	return fees, nil
}

func StadiumFeeCalculator() *stadiumFeeCalculator {
	fees := models.NewParkingLotFees()
	return &stadiumFeeCalculator{
		*fees,
	}
}

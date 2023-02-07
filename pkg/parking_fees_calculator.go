package pkg

import (
	"fmt"
	"math"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

func parkingRateFor(durationKey string, parkingRates map[string]int64) (int64, error) {
	if parkingRate, ok := parkingRates[durationKey]; ok {
		return parkingRate, nil
	}
	return 0, fmt.Errorf("error: invalid Key %s", durationKey)
}

func calculateParkingFees(
	vehicleType common.VehicleType,
	feeDetails models.ParkingLotFeeDetails,
	totalTimeParked time.Duration,
) (*common.Currency, error) {
	var feeAmount int64 = 0
	parkingRates, err := feeDetails.ParkingRates(vehicleType)
	if err != nil {
		return nil, err
	}
	switch feeDetails.FeesModelName() {
	case common.Mall:
		perHourRateInSub, err := parkingRateFor(common.KeyPerHourRate, parkingRates)
		if err != nil {
			return nil, err
		}
		chargeableHours := int64(math.Ceil(totalTimeParked.Hours()))
		feeAmount = chargeableHours * perHourRateInSub
	case common.Airport:
		chargeableHours := totalTimeParked.Hours()
		if chargeableHours >= 24 {
			chargeableDays := int64(math.Ceil(chargeableHours / 24))
			feeAmount = chargeableDays * parkingRates[common.KeyPerDay]
		} else {
			if vehicleType == common.TwoWheeler {
				if chargeableHours >= 8 {
					feeAmount = parkingRates[common.Key_8_24_Hrs]
				} else if chargeableHours >= 1 {
					feeAmount = parkingRates[common.Key_1_8_Hrs]
				} else {
					feeAmount = parkingRates[common.Key_0_1_Hrs]
				}
			} else {
				if chargeableHours >= 12 {
					feeAmount = parkingRates[common.Key_12_24_Hrs]
				} else {
					feeAmount = parkingRates[common.Key_0_12_Hrs]
				}
			}
		}
	case common.Stadium:
		chargeableHours := totalTimeParked.Hours()
		feeAmount += parkingRates[common.Key_0_4_Hrs]
		if chargeableHours >= 4 {
			feeAmount += parkingRates[common.Key_4_12_Hrs]
		}
		if chargeableHours >= 12 {
			chargeableHoursAfter12 := int64(math.Ceil(chargeableHours - (11 + ((59.0 / 60) * (1 + (0.1 / 6))))))
			feeAmount += (parkingRates[common.Key_12_Inf_PerHr] * chargeableHoursAfter12)
		}
	default:
		return nil, fmt.Errorf("error: %s parking fee details not available", feeDetails.FeesModelName())
	}
	feeAmountInSubUnit := feeAmount * common.DefaultCurrencySubUnitConversionFactor
	parkingFees, err := common.NewCurrency(feeAmountInSubUnit, common.DefaultCurrencyUnit, common.DefaultCurrencySubUnitConversionFactor)
	if err != nil {
		return nil, err
	}
	return parkingFees, nil
}

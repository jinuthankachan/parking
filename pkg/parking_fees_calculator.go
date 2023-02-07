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
	}
	feeAmountInSubUnit := feeAmount * common.DefaultCurrencySubUnitConversionFactor
	parkingFees, err := common.NewCurrency(feeAmountInSubUnit, common.DefaultCurrencyUnit, common.DefaultCurrencySubUnitConversionFactor)
	if err != nil {
		return nil, err
	}
	return parkingFees, nil
}

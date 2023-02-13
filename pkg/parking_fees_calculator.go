package pkg

import (
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type ParkingFeeCalculator interface {
	models.ParkingRatesFetcher
	Calculate(
		vehicleType common.VehicleType,
		totalTimeParked time.Duration,
	) (*common.Currency, error)
}

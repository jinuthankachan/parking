package models

import (
	"strconv"

	"github.com/jinut2/parking/common"
)

type ParkingLotModel struct {
	FeesModelName common.FeesModelName
	Spots         map[string]SpotModel
}

type SpotModel struct {
	Type     common.VehicleType
	Occupied bool
}

type SpotRegister interface {
	AssignSpot(common.VehicleType) (spotID string, err error)
	SpotDetails(spotID string) (*SpotModel, error)
	UnassignSpot(spotID string) error
}

type ParkingLot struct {
	ParkingLotModel
}

func NewParkingLot(
	feesModelName common.FeesModelName,
	TwoWheeleroWheelerSpots int64,
	fourWheelLightVehicleSpots int64,
	heavyVehicleSpots int64,
) *ParkingLot {
	totalSpots := TwoWheeleroWheelerSpots + fourWheelLightVehicleSpots + heavyVehicleSpots
	spots := make(map[string]SpotModel, totalSpots)
	spotIdx := 1
	for i := 0; i < int(TwoWheeleroWheelerSpots); i++ {
		spot := SpotModel{
			Type:     common.TwoWheeler,
			Occupied: false,
		}
		spots[strconv.Itoa(spotIdx)] = spot
		spotIdx++
	}
	for i := 0; i < int(fourWheelLightVehicleSpots); i++ {
		spot := SpotModel{
			Type:     common.Light4Wheeler,
			Occupied: false,
		}
		spots[strconv.Itoa(spotIdx)] = spot
		spotIdx++
	}
	for i := 0; i < int(heavyVehicleSpots); i++ {
		spot := SpotModel{
			Type:     common.HeavyVehicle,
			Occupied: false,
		}
		spots[strconv.Itoa(spotIdx)] = spot
		spotIdx++
	}
	return &ParkingLot{
		ParkingLotModel: ParkingLotModel{
			FeesModelName: feesModelName,
			Spots:         spots,
		},
	}
}

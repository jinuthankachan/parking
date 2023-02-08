package models

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/jinut2/parking/common"
)

type ParkingLotModel struct {
	FeesModelName common.FeesModelName
	Spots         []*SpotModel
}

type SpotModel struct {
	ID       string
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
	mu sync.Mutex
}

func NewParkingLot(
	feesModelName common.FeesModelName,
	TwoWheeleroWheelerSpots int64,
	fourWheelLightVehicleSpots int64,
	heavyVehicleSpots int64,
) *ParkingLot {
	spots := []*SpotModel{}
	spotIdx := 1
	for i := 0; i < int(TwoWheeleroWheelerSpots); i++ {
		spot := SpotModel{
			ID:       strconv.Itoa(spotIdx),
			Type:     common.TwoWheeler,
			Occupied: false,
		}
		spots = append(spots, &spot)
		spotIdx++
	}
	for i := 0; i < int(fourWheelLightVehicleSpots); i++ {
		spot := SpotModel{
			ID:       strconv.Itoa(spotIdx),
			Type:     common.Light4Wheeler,
			Occupied: false,
		}
		spots = append(spots, &spot)
		spotIdx++
	}
	for i := 0; i < int(heavyVehicleSpots); i++ {
		spot := SpotModel{
			ID:       strconv.Itoa(spotIdx),
			Type:     common.HeavyVehicle,
			Occupied: false,
		}
		spots = append(spots, &spot)
		spotIdx++
	}
	return &ParkingLot{
		ParkingLotModel: ParkingLotModel{
			FeesModelName: feesModelName,
			Spots:         spots,
		},
	}
}

func (pl *ParkingLot) AssignSpot(vehicleType common.VehicleType) (spotID string, err error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	for id, spot := range pl.Spots {
		if spot.Type == vehicleType && !spot.Occupied {
			spot.Occupied = true
			pl.Spots[id] = spot
			return spot.ID, nil
		}
	}
	return "", fmt.Errorf("error: No spots for %s is available", vehicleType)
}
func (pl *ParkingLot) SpotDetails(spotID string) (*SpotModel, error) {
	for _, spot := range pl.Spots {
		if spot.ID == spotID {
			return spot, nil
		}
	}
	return nil, fmt.Errorf("error: spot %s not found", spotID)
}
func (pl *ParkingLot) UnassignSpot(spotID string) error {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	for i, spot := range pl.Spots {
		if spot.ID == spotID {
			spot.Occupied = false
			pl.Spots[i] = spot
			return nil
		}
	}
	return fmt.Errorf("error: spot %s not found", spotID)
}

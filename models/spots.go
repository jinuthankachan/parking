package models

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/jinut2/parking/common"
)

type SpotModel struct {
	Type     common.VehicleType
	Occupied bool
}

type SpotRegister interface {
	AssignSpot(common.VehicleType) (spotID string, err error)
	SpotDetails(spotID string) (*SpotModel, error)
	UnassignSpot(spotID string) error
}

type Spots struct {
	mu    sync.Mutex
	store map[string]SpotModel
}

func NewSpotsRegister(
	twoWheeleroWheelerSpots int64,
	fourWheelLightVehicleSpots int64,
	heavyVehicleSpots int64,
) *Spots {
	totalSpots := twoWheeleroWheelerSpots + fourWheelLightVehicleSpots + heavyVehicleSpots
	spots := make(map[string]SpotModel, totalSpots)
	spotIdx := 1
	for i := 0; i < int(twoWheeleroWheelerSpots); i++ {
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
	return &Spots{
		store: spots,
	}
}

func (s *Spots) AssignSpot(vehicleType common.VehicleType) (spotID string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := 0; i < len(s.store); i++ {
		id := strconv.Itoa(i)
		spot := s.store[id]
		if spot.Type == vehicleType && !spot.Occupied {
			spot.Occupied = true
			s.store[id] = spot
			return id, nil
		}
	}
	return "", fmt.Errorf("error: No spots for %s is available", vehicleType)
}
func (s *Spots) SpotDetails(spotID string) (*SpotModel, error) {
	if spot, ok := s.store[spotID]; ok {
		return &spot, nil
	}
	return nil, fmt.Errorf("error: spot %s not found", spotID)
}
func (s *Spots) UnassignSpot(spotID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if spot, ok := s.store[spotID]; ok {
		spot.Occupied = false
		s.store[spotID] = spot
		return nil
	}
	return fmt.Errorf("error: spot %s not found", spotID)
}

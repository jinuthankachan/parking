package models

import (
	"fmt"
	"math"
	"time"

	"github.com/jinut2/parking/common"
)

type ParkingLotFeeModel struct {
	FeesModelName common.FeesModelName
	VehicleType   common.VehicleType
	SlabType      common.FeeSlabType
	LowerLimit    time.Duration
	UpperLimit    time.Duration
	Charge        *common.Currency
}

type ParkingLotFees struct {
	store []ParkingLotFeeModel
}

type ParkingRatesFetcher interface {
	GetAllSlabsUnderTime(upperLimit time.Duration, feeModelName common.FeesModelName, vehicleType common.VehicleType) ([]ParkingLotFeeModel, error)
	GetSlabForTime(duration time.Duration, feeModelName common.FeesModelName, vehicleType common.VehicleType) (*ParkingLotFeeModel, error)
}

func NewParkingLotFees() *ParkingLotFees {
	store := []ParkingLotFeeModel{
		{
			FeesModelName: common.Mall,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 10),
		},
		{
			FeesModelName: common.Mall,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 20),
		},
		{
			FeesModelName: common.Mall,
			VehicleType:   common.HeavyVehicle,
			SlabType:      common.FlatSlab,
			LowerLimit:    0,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 50),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0,
			UpperLimit:    1 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 0),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    1 * time.Hour,
			UpperLimit:    8 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 40),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    8 * time.Hour,
			UpperLimit:    24 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 60),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.PerDaySlab,
			LowerLimit:    24 * time.Hour,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 80),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0 * time.Hour,
			UpperLimit:    12 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 60),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    12 * time.Hour,
			UpperLimit:    24 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 80),
		},
		{
			FeesModelName: common.Airport,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.PerDaySlab,
			LowerLimit:    24 * time.Hour,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 100),
		},
		{
			FeesModelName: common.Stadium,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0 * time.Hour,
			UpperLimit:    4 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 30),
		},
		{
			FeesModelName: common.Stadium,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    4 * time.Hour,
			UpperLimit:    12 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 60),
		},
		{
			FeesModelName: common.Stadium,
			VehicleType:   common.TwoWheeler,
			SlabType:      common.PerHourSlab,
			LowerLimit:    12 * time.Hour,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 100),
		},

		{
			FeesModelName: common.Stadium,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    0 * time.Hour,
			UpperLimit:    4 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 60),
		},
		{
			FeesModelName: common.Stadium,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.FlatSlab,
			LowerLimit:    4 * time.Hour,
			UpperLimit:    12 * time.Hour,
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 120),
		},
		{
			FeesModelName: common.Stadium,
			VehicleType:   common.Light4Wheeler,
			SlabType:      common.PerHourSlab,
			LowerLimit:    12 * time.Hour,
			UpperLimit:    time.Duration(math.MaxInt64),
			Charge:        common.DefaultCurrency(common.DefaultCurrencySubUnitConversionFactor * 200),
		},
	}

	return &ParkingLotFees{
		store: store,
	}
}

func (plf *ParkingLotFees) GetAllSlabsUnderTime(duration time.Duration,
	feeModelName common.FeesModelName,
	vehicleType common.VehicleType,
) (slabs []ParkingLotFeeModel, err error) {
	for _, slab := range plf.store {
		if slab.FeesModelName == feeModelName && vehicleType == slab.VehicleType && (duration >= slab.LowerLimit) {
			slabs = append(slabs, slab)
		}
	}
	if len(slabs) == 0 {
		err = fmt.Errorf("error: no slabs found for %+v", duration)
	}
	return
}

func (plf *ParkingLotFees) GetSlabForTime(duration time.Duration,
	feeModelName common.FeesModelName,
	vehicleType common.VehicleType,
) (*ParkingLotFeeModel, error) {
	for _, slab := range plf.store {
		if slab.FeesModelName == feeModelName && vehicleType == slab.VehicleType && (duration < slab.UpperLimit) && (duration >= slab.LowerLimit) {
			return &slab, nil
		}
	}
	return nil, fmt.Errorf("error: no slabs found for %+v", duration)
}

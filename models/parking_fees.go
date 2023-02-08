package models

import (
	"fmt"

	"github.com/jinut2/parking/common"
)

type ParkingLotFeeDetails interface {
	FeesModelName() common.FeesModelName
	ParkingRates(common.VehicleType) (map[string]int64, error)
}

type ParkingLotFeesModel struct {
	feesModelName common.FeesModelName
	tariff        map[common.VehicleType]map[string]int64
}

func NewParkingLotFeesModel(feesModelName common.FeesModelName) *ParkingLotFeesModel {
	tariff := map[common.FeesModelName]map[common.VehicleType]map[string]int64{
		common.Mall: {
			common.TwoWheeler: {
				common.KeyPerHourRate: 10,
			},
			common.Light4Wheeler: {
				common.KeyPerHourRate: 20,
			},
			common.HeavyVehicle: {
				common.KeyPerHourRate: 50,
			},
		},
		common.Airport: {
			common.TwoWheeler: {
				common.Key_0_1_Hrs:  0,
				common.Key_1_8_Hrs:  40,
				common.Key_8_24_Hrs: 60,
				common.KeyPerDay:    80,
			},
			common.Light4Wheeler: {
				common.Key_0_12_Hrs:  60,
				common.Key_12_24_Hrs: 80,
				common.KeyPerDay:     100,
			},
		},
		common.Stadium: {
			common.TwoWheeler: {
				common.Key_0_4_Hrs:      30,
				common.Key_4_12_Hrs:     60,
				common.Key_12_Inf_PerHr: 100,
			},
			common.Light4Wheeler: {
				common.Key_0_4_Hrs:      60,
				common.Key_4_12_Hrs:     120,
				common.Key_12_Inf_PerHr: 200,
			},
		},
	}
	return &ParkingLotFeesModel{
		feesModelName: feesModelName,
		tariff:        tariff[feesModelName],
	}
}

func (fm *ParkingLotFeesModel) ParkingRates(vehicleType common.VehicleType) (map[string]int64, error) {
	if tariff, ok := fm.tariff[vehicleType]; ok {
		return tariff, nil
	}
	return nil, fmt.Errorf("error: %s vehicle not allowed in %s parking", vehicleType, fm.feesModelName)
}

func (fm *ParkingLotFeesModel) FeesModelName() common.FeesModelName {
	return fm.feesModelName
}

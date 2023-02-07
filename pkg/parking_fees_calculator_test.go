package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

func defaultCurrency(valueInSubUnit int64) *common.Currency {
	m, _ := common.NewCurrency(valueInSubUnit,
		common.DefaultCurrencyUnit,
		common.DefaultCurrencySubUnitConversionFactor)
	return m
}

func Test_calculateParkingFees(t *testing.T) {
	type args struct {
		vehicleType     common.VehicleType
		feesModelName   common.FeesModelName
		totalTimeParked time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *common.Currency
		wantErr bool
	}{
		{
			name: "Park motorcycle in Mall for 3 hours 30 mins",
			args: args{
				common.TwoWheeler,
				common.Mall,
				(3 * time.Hour) + (30 * time.Minute),
			},
			want:    defaultCurrency(4000),
			wantErr: false,
		},
		{
			name: "Park bike in Mall for 0 hours 0 min",
			args: args{
				common.Light4Wheeler,
				common.Mall,
				(0 * time.Hour) + (0 * time.Minute),
			},
			want:    defaultCurrency(0),
			wantErr: false,
		},
		{
			name: "Park car in Mall for 6 hours 01 min",
			args: args{
				common.Light4Wheeler,
				common.Mall,
				(6 * time.Hour) + (1 * time.Minute),
			},
			want:    defaultCurrency(14000),
			wantErr: false,
		},
		{
			name: "Park truck in Mall for 1 hours 59 mins",
			args: args{
				common.HeavyVehicle,
				common.Mall,
				(1 * time.Hour) + (59 * time.Minute),
			},
			want:    defaultCurrency(10000),
			wantErr: false,
		},
		{
			name: "Park 2 wheeler in Stadium for 3 hours 40 mins",
			args: args{
				common.TwoWheeler,
				common.Stadium,
				(3 * time.Hour) + (40 * time.Minute),
			},
			want:    defaultCurrency(3000),
			wantErr: false,
		},
		{
			name: "Park 2 wheeler in Stadium for 14 hours 59 mins",
			args: args{
				common.TwoWheeler,
				common.Stadium,
				(14 * time.Hour) + (59 * time.Minute),
			},
			want:    defaultCurrency(39000),
			wantErr: false,
		},
		{
			name: "Park SUV in Stadium for 3 hours 40 mins",
			args: args{
				common.Light4Wheeler,
				common.Stadium,
				(11 * time.Hour) + (30 * time.Minute),
			},
			want:    defaultCurrency(18000),
			wantErr: false,
		},
		{
			name: "Park SUV in Stadium for 13 hours 5 mins",
			args: args{
				common.Light4Wheeler,
				common.Stadium,
				(13 * time.Hour) + (5 * time.Minute),
			},
			want:    defaultCurrency(58000),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plfm := models.NewParkingLotFeesModel(tt.args.feesModelName)
			got, err := calculateParkingFees(tt.args.vehicleType, plfm, tt.args.totalTimeParked)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateParkingFees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateParkingFees() = %v, want %v", got, tt.want)
			}
		})
	}
}

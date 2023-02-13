package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
)

func Test_mallFeeCalculator_Calculate(t *testing.T) {
	type args struct {
		vehicleType     common.VehicleType
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
				(3 * time.Hour) + (30 * time.Minute),
			},
			want:    common.DefaultCurrency(4000),
			wantErr: false,
		},
		{
			name: "Park bike in Mall for 0 hours 0 min",
			args: args{
				common.Light4Wheeler,
				(0 * time.Hour) + (0 * time.Minute),
			},
			want:    common.DefaultCurrency(0),
			wantErr: false,
		},
		{
			name: "Park car in Mall for 6 hours 01 min",
			args: args{
				common.Light4Wheeler,
				(6 * time.Hour) + (1 * time.Minute),
			},
			want:    common.DefaultCurrency(14000),
			wantErr: false,
		},
		{
			name: "Park truck in Mall for 1 hours 59 mins",
			args: args{
				common.HeavyVehicle,
				(1 * time.Hour) + (59 * time.Minute),
			},
			want:    common.DefaultCurrency(10000),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := MallFeeCalculator()
			got, err := calc.Calculate(tt.args.vehicleType, tt.args.totalTimeParked)
			if (err != nil) != tt.wantErr {
				t.Errorf("mallFeeCalculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mallFeeCalculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

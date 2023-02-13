package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
)

func Test_stadiumFeeCalculator_Calculate(t *testing.T) {
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
			name: "Park 2 wheeler in Stadium for 3 hours 40 mins",
			args: args{
				common.TwoWheeler,
				(3 * time.Hour) + (40 * time.Minute),
			},
			want:    common.DefaultCurrency(3000),
			wantErr: false,
		},
		{
			name: "Park 2 wheeler in Stadium for 14 hours 59 mins",
			args: args{
				common.TwoWheeler,
				(14 * time.Hour) + (59 * time.Minute),
			},
			want:    common.DefaultCurrency(39000),
			wantErr: false,
		},
		{
			name: "Park SUV in Stadium for 3 hours 40 mins",
			args: args{
				common.Light4Wheeler,
				(11 * time.Hour) + (30 * time.Minute),
			},
			want:    common.DefaultCurrency(18000),
			wantErr: false,
		},
		{
			name: "Park SUV in Stadium for 13 hours 5 mins",
			args: args{
				common.Light4Wheeler,
				(13 * time.Hour) + (5 * time.Minute),
			},
			want:    common.DefaultCurrency(58000),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := StadiumFeeCalculator()
			got, err := calc.Calculate(tt.args.vehicleType, tt.args.totalTimeParked)
			if (err != nil) != tt.wantErr {
				t.Errorf("stadiumFeeCalculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stadiumFeeCalculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

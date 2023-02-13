package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
)

func Test_airportFeeCalculator_Calculate(t *testing.T) {
	type fields struct{}
	type args struct {
		vehicleType     common.VehicleType
		totalTimeParked time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Currency
		wantErr bool
	}{
		{
			name: "Park 2 wheeler in Airport for 55 mins",
			args: args{
				common.TwoWheeler,
				(0 * time.Hour) + (55 * time.Minute),
			},
			want:    common.DefaultCurrency(0),
			wantErr: false,
		},
		{
			name: "Park 2 wheeler in Airport for 14 hrs 59 mins",
			args: args{
				common.TwoWheeler,
				(14 * time.Hour) + (59 * time.Minute),
			},
			want:    common.DefaultCurrency(6000),
			wantErr: false,
		},
		{
			name: "Park 2 wheeler in Airport for 1 day 12 hrs",
			args: args{
				common.TwoWheeler,
				((24 + 12) * time.Hour),
			},
			want:    common.DefaultCurrency(16000),
			wantErr: false,
		},
		{
			name: "Park car in Airport for 55 mins",
			args: args{
				common.Light4Wheeler,
				(0 * time.Hour) + (55 * time.Minute),
			},
			want:    common.DefaultCurrency(6000),
			wantErr: false,
		},
		{
			name: "Park car in Airport for 23 hrs 59 mins",
			args: args{
				common.Light4Wheeler,
				(23 * time.Hour) + (59 * time.Minute),
			},
			want:    common.DefaultCurrency(8000),
			wantErr: false,
		},
		{
			name: "Park car in Airport for 3 days 1 hrs",
			args: args{
				common.Light4Wheeler,
				(((3 * 24) + 1) * time.Hour),
			},
			want:    common.DefaultCurrency(40000),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := AirportFeeCalculator()
			got, err := calc.Calculate(tt.args.vehicleType, tt.args.totalTimeParked)
			if (err != nil) != tt.wantErr {
				t.Errorf("airportFeeCalculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("airportFeeCalculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

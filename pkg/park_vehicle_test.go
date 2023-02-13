package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

var timeZone *time.Location
var mallSpotsRegister *models.Spots
var mallTicketCounter *models.Tickets

func init() {
	var err error
	timeZone, err = time.LoadLocation(common.DefaultTimezone)
	if err != nil {
		panic(err)
	}
	mallSpotsRegister = models.NewSpotsRegister(3, 2, 1)
	mallTicketCounter = models.NewTicketCounter()
}

func TestParkVehicle(t *testing.T) {

	type args struct {
		spotRegister  models.SpotRegister
		ticketCounter models.TicketCounter
		vehicleType   common.VehicleType
		entryTime     string
		timeZone      *time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    *ParkingTicket
		wantErr bool
	}{
		{
			name: "Park a bike 01",
			args: args{
				spotRegister:  mallSpotsRegister,
				ticketCounter: mallTicketCounter,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 14:04:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketID:  "1",
				SpotID:    "1",
				EntryTime: "29-May-2022 14:04:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter 02",
			args: args{
				spotRegister:  mallSpotsRegister,
				ticketCounter: mallTicketCounter,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 14:44:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketID:  "2",
				SpotID:    "2",
				EntryTime: "29-May-2022 14:44:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter 03",
			args: args{
				spotRegister:  mallSpotsRegister,
				ticketCounter: mallTicketCounter,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 15:59:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketID:  "3",
				SpotID:    "3",
				EntryTime: "29-May-2022 15:59:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter 04",
			args: args{
				spotRegister:  mallSpotsRegister,
				ticketCounter: mallTicketCounter,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 16:54:07",
				timeZone:      timeZone,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Park a car 01",
			args: args{
				spotRegister:  mallSpotsRegister,
				ticketCounter: mallTicketCounter,
				vehicleType:   common.Light4Wheeler,
				entryTime:     "29-May-2022 13:54:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketID:  "4",
				SpotID:    "4",
				EntryTime: "29-May-2022 13:54:07",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vehicleEntry := models.MockVehicleEntry(tt.args.vehicleType, tt.args.entryTime, tt.args.timeZone)
			got, err := ParkVehicle(tt.args.spotRegister, tt.args.ticketCounter, vehicleEntry, tt.args.timeZone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

package pkg

import (
	"reflect"
	"testing"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

var timeZone = common.DefaultTimezone
var mallParkingLot = models.NewParkingLot(common.Mall, 3, 2, 1)
var mallTickets = models.NewTickets()

func TestParkVehicle(t *testing.T) {

	type args struct {
		spotRegister  models.SpotRegister
		ticketCounter models.TicketCounter
		vehicleType   common.VehicleType
		entryTime     string
		timeZone      string
	}
	tests := []struct {
		name    string
		args    args
		want    *ParkingTicket
		wantErr bool
	}{
		{
			name: "Park a bike",
			args: args{
				spotRegister:  mallParkingLot,
				ticketCounter: mallTickets,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 14:04:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketNumber: "1",
				SpotNumber:   "1",
				EntryTime:    "29-May-2022 14:04:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter",
			args: args{
				spotRegister:  mallParkingLot,
				ticketCounter: mallTickets,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 14:44:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketNumber: "2",
				SpotNumber:   "2",
				EntryTime:    "29-May-2022 14:44:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter",
			args: args{
				spotRegister:  mallParkingLot,
				ticketCounter: mallTickets,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 15:54:07",
				timeZone:      timeZone,
			},
			want: &ParkingTicket{
				TicketNumber: "3",
				SpotNumber:   "3",
				EntryTime:    "29-May-2022 15:54:07",
			},
			wantErr: false,
		},
		{
			name: "Park a scooter",
			args: args{
				spotRegister:  mallParkingLot,
				ticketCounter: mallTickets,
				vehicleType:   common.TwoWheeler,
				entryTime:     "29-May-2022 16:54:07",
				timeZone:      timeZone,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParkVehicle(tt.args.spotRegister, tt.args.ticketCounter, tt.args.vehicleType, tt.args.entryTime, tt.args.timeZone)
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

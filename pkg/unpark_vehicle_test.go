package pkg

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

var mallReceiptBook = models.NewReceiptBook()
var mallFeeDetails = models.NewParkingLotFeesModel(common.Mall)

func TestUnparkVehicle(t *testing.T) {
	type args struct {
		ticketID             string
		exitTime             string
		timeZone             *time.Location
		spotRegister         models.SpotRegister
		ticketCounter        models.TicketCounter
		parkingLotFeeDetails models.ParkingLotFeeDetails
		ReceiptBook          models.ReceiptBook
	}
	tests := []struct {
		name    string
		args    args
		want    *ParkingReceipt
		wantErr bool
	}{
		{
			name: "Unpark scooter 02",
			args: args{
				ticketID:             "2",
				exitTime:             "29-May-2022 15:40:07",
				timeZone:             timeZone,
				spotRegister:         mallSpotsRegister,
				ticketCounter:        mallTicketCounter,
				parkingLotFeeDetails: mallFeeDetails,
				ReceiptBook:          mallReceiptBook,
			},
			want: &ParkingReceipt{
				ReceiptID: "R-1",
				EntryTime: "29-May-2022 14:44:07",
				ExitTime:  "29-May-2022 15:40:07",
				Fees:      "₹ 10.00",
			},
			wantErr: false,
		},
		{
			name: "Unpark scooter 01",
			args: args{
				ticketID:             "1",
				exitTime:             "29-May-2022 17:44:07",
				timeZone:             timeZone,
				spotRegister:         mallSpotsRegister,
				ticketCounter:        mallTicketCounter,
				parkingLotFeeDetails: mallFeeDetails,
				ReceiptBook:          mallReceiptBook,
			},
			want: &ParkingReceipt{
				ReceiptID: "R-2",
				EntryTime: "29-May-2022 14:04:07",
				ExitTime:  "29-May-2022 17:44:07",
				Fees:      "₹ 40.00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vehicleExit := models.MockVehicleExit(tt.args.ticketID, tt.args.exitTime, tt.args.timeZone)
			got, err := UnparkVehicle(vehicleExit, tt.args.spotRegister, tt.args.ticketCounter, tt.args.parkingLotFeeDetails, tt.args.ReceiptBook, tt.args.timeZone)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnparkVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnparkVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

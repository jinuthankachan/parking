package pkg

import (
	"reflect"
	"testing"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

var mallReceipts = models.NewReceipts()
var mallFeeDetails = models.NewParkingLotFeesModel(common.Mall)

func TestUnparkVehicle(t *testing.T) {
	type args struct {
		ticketID             string
		exitTime             string
		timeZone             string
		spotRegister         models.SpotRegister
		ticketCounter        models.TicketCounter
		parkingLotFeeDetails models.ParkingLotFeeDetails
		receiptGenerator     models.ReceiptGenerator
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
				spotRegister:         mallParkingLot,
				ticketCounter:        mallTickets,
				parkingLotFeeDetails: mallFeeDetails,
				receiptGenerator:     mallReceipts,
			},
			want: &ParkingReceipt{
				ReceiptNumber: "R-1",
				EntryTime:     "29-May-2022 14:44:07",
				ExitTime:      "29-May-2022 15:40:07",
				Fees:          "₹ 10.00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnparkVehicle(tt.args.ticketID, tt.args.exitTime, tt.args.timeZone, tt.args.spotRegister, tt.args.ticketCounter, tt.args.parkingLotFeeDetails, tt.args.receiptGenerator)
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

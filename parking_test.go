package parking

import (
	"testing"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

var timeZone = common.DefaultTimezone

func TestMallParkingLot(t *testing.T) {
	var mallParkingLot = models.NewParkingLot(common.Mall, 3, 2, 1)
	var mallTickets = models.NewTickets()
	var mallReceipts = models.NewReceipts()
	var mallFeeDetails = models.NewParkingLotFeesModel(common.Mall)

	t.Run(
		"Park 2W-01",
		func(t *testing.T) {
			got, err := pkg.ParkVehicle(mallParkingLot, mallTickets, common.TwoWheeler, tt.args.entryTime, tt.args.timeZone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParkVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParkVehicle() = %v, want %v", got, tt.want)
			}
		})
	)
}

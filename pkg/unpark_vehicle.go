package pkg

import (
	"fmt"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type ParkingReceipt struct {
	ReceiptID string
	EntryTime string
	ExitTime  string
	Fees      string
}

func (r *ParkingReceipt) Print() string {
	return fmt.Sprintf("%+v", *r)
}

func UnparkVehicle(
	vehicleExit models.VehicleExit,
	spotRegister models.SpotRegister,
	ticketCounter models.TicketCounter,
	parkingLotFeeDetails models.ParkingLotFeeDetails,
	ReceiptBook models.ReceiptBook,
	timeZone *time.Location,
) (*ParkingReceipt, error) {
	exitTime := vehicleExit.ExitTime()
	ticketDetails, err := ticketCounter.TicketDetails(vehicleExit.TicketID())
	if err != nil {
		return nil, err
	}
	totalTimeParked := exitTime.Sub(ticketDetails.EntryTime)
	spotDetails, err := spotRegister.SpotDetails(ticketDetails.SpotID)
	if err != nil {
		return nil, err
	}
	parkingFees, err := calculateParkingFees(
		spotDetails.Type,
		parkingLotFeeDetails,
		totalTimeParked,
	)
	if err != nil {
		return nil, err
	}
	receiptID, err := ReceiptBook.GenerateReceipt(ticketDetails, parkingFees, exitTime)
	if err != nil {
		return nil, err
	}
	return &ParkingReceipt{
		ReceiptID: receiptID,
		EntryTime: ticketDetails.EntryTime.In(timeZone).Format(common.DefaultTimeFormat),
		ExitTime:  exitTime.In(timeZone).Format(common.DefaultTimeFormat),
		Fees:      parkingFees.DisplayValue(),
	}, nil
}

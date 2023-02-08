package pkg

import (
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

type ParkingReceipt struct {
	ReceiptNumber string
	EntryTime     string
	ExitTime      string
	Fees          string
}

func UnparkVehicle(
	ticketID string,
	exitTime string,
	timeZone string,
	spotRegister models.SpotRegister,
	ticketCounter models.TicketCounter,
	parkingLotFeeDetails models.ParkingLotFeeDetails,
	receiptGenerator models.ReceiptGenerator,
) (*ParkingReceipt, error) {
	if timeZone == "" {
		timeZone = common.DefaultTimezone
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}
	parsedExitTime, err := time.ParseInLocation(common.DefaultTimeFormat, exitTime, loc)
	if err != nil {
		panic(err)
	}
	ticketDetails, err := ticketCounter.TicketDetails(ticketID)
	if err != nil {
		return nil, err
	}
	totalTimeParked := parsedExitTime.Sub(ticketDetails.EntryTime)
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
	receiptID, err := receiptGenerator.GenerateReceipt(ticketDetails, parkingFees, parsedExitTime)
	if err != nil {
		return nil, err
	}
	return &ParkingReceipt{
		ReceiptNumber: receiptID,
		EntryTime:     ticketDetails.EntryTime.In(loc).Format(common.DefaultTimeFormat),
		ExitTime:      exitTime,
		Fees:          parkingFees.DisplayValue(),
	}, nil
}

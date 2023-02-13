package services

import (
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
	"github.com/jinut2/parking/pkg"
)

// ParkingLotService : This provides the methods required for parking and unparking vehicles.
//
//	A new instance can be created using NewParkingLot method
type ParkingLotService struct {
	spots          *models.Spots
	tickets        *models.Tickets
	receipts       *models.Receipts
	feesCalculator pkg.ParkingFeeCalculator
	timeZone       *time.Location
}

func NewParkingLot(
	timeZone string,
	twoWheelerSpots int64,
	light4WheelerSpots int64,
	heavyVehicleSpots int64,
	feeCalculator pkg.ParkingFeeCalculator,
) (*ParkingLotService, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}

	spotRegister := models.NewSpotsRegister(twoWheelerSpots, light4WheelerSpots, heavyVehicleSpots)
	tickets := models.NewTicketCounter()
	receipts := models.NewReceiptBook()
	return &ParkingLotService{
		spots:          spotRegister,
		tickets:        tickets,
		receipts:       receipts,
		feesCalculator: feeCalculator,
		timeZone:       loc,
	}, nil
}

func (pl *ParkingLotService) ParkVehicle(vehicleType common.VehicleType) (*pkg.ParkingTicket, error) {
	return pkg.ParkVehicle(pl.spots, pl.tickets, models.NewVehicleEntry(vehicleType), pl.timeZone)
}

func (pl *ParkingLotService) UnparkVehicle(ticketID string) (*pkg.ParkingReceipt, error) {
	return pkg.UnparkVehicle(models.NewVehicleExit(ticketID), pl.spots, pl.tickets, pl.receipts, pl.feesCalculator, pl.timeZone)
}

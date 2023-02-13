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
	spots         *models.Spots
	tickets       *models.Tickets
	receipts      *models.Receipts
	feeDetails    *models.ParkingLotFeesModel
	feesModelName common.FeesModelName
	timeZone      *time.Location
}

func NewParkingLot(
	parkingLotType, timeZone string,
	twoWheelerSpots int64,
	light4WheelerSpots int64,
	heavyVehicleSpots int64,
) (*ParkingLotService, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}

	feeModelName, err := common.FeesModelNameFromParkingType(parkingLotType)
	if err != nil {
		return nil, err
	}

	spotRegister := models.NewSpotsRegister(twoWheelerSpots, light4WheelerSpots, heavyVehicleSpots)
	tickets := models.NewTicketCounter()
	feesModel := models.NewParkingLotFeesModel(feeModelName)
	receipts := models.NewReceiptBook()
	return &ParkingLotService{
		spots:         spotRegister,
		tickets:       tickets,
		feeDetails:    feesModel,
		receipts:      receipts,
		feesModelName: feeModelName,
		timeZone:      loc,
	}, nil
}

func (pl *ParkingLotService) ParkVehicle(vehicleType common.VehicleType) (*pkg.ParkingTicket, error) {
	return pkg.ParkVehicle(pl.spots, pl.tickets, models.NewVehicleEntry(vehicleType), pl.timeZone)
}

func (pl *ParkingLotService) UnparkVehicle(ticketID string) (*pkg.ParkingReceipt, error) {
	return pkg.UnparkVehicle(models.NewVehicleExit(ticketID), pl.spots, pl.tickets, pl.feeDetails, pl.receipts, pl.timeZone)
}

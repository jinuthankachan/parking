package pkg

import (
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/internal/models"
)

type ParkingTicket struct {
	TicketNumber string
	SpotNumber   string
	EntryTime    string
}

func ParkVehicle(
	spotRegister models.SpotRegister,
	ticketCounter models.TicketCounter,
	vehicleType common.VehicleType,
	entryTime string,
	timeZone string,
) (*ParkingTicket, error) {
	if timeZone == "" {
		timeZone = common.DefaultTimezone
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}
	parsedEntryTime, err := time.ParseInLocation(common.DefaultTimeFormat, entryTime, loc)
	if err != nil {
		return nil, err
	}
	reservedSpotID, err := spotRegister.AssignSpot(vehicleType)
	if err != nil {
		return nil, err
	}
	allotedTicketID, err := ticketCounter.AllotTicket(reservedSpotID, parsedEntryTime)
	if err != nil {
		return nil, err
	}
	return &ParkingTicket{
		TicketNumber: allotedTicketID,
		SpotNumber:   reservedSpotID,
		EntryTime:    entryTime,
	}, nil
}

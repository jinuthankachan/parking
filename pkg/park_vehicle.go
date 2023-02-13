package pkg

import (
	"fmt"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
)

type ParkingTicket struct {
	TicketID  string
	SpotID    string
	EntryTime string
}

func (r *ParkingTicket) Print() string {
	return fmt.Sprintf("%+v", *r)
}

func ParkVehicle(
	spotRegister models.SpotRegister,
	ticketCounter models.TicketCounter,
	vehicleEntry models.VehicleEntry,
	timeZone *time.Location,
) (*ParkingTicket, error) {
	reservedSpotID, err := spotRegister.AssignSpot(vehicleEntry.VehicleType())
	if err != nil {
		return nil, err
	}
	vehicleEntryTime := vehicleEntry.EntryTime()
	allotedTicketID, err := ticketCounter.AllotTicket(reservedSpotID, vehicleEntryTime)
	if err != nil {
		return nil, err
	}
	return &ParkingTicket{
		TicketID:  allotedTicketID,
		SpotID:    reservedSpotID,
		EntryTime: vehicleEntryTime.In(timeZone).Format(common.DefaultTimeFormat),
	}, nil
}

package models

import (
	"time"

	"github.com/jinut2/parking/common"
)

type VehicleEntry interface {
	VehicleType() common.VehicleType
	EntryTime() time.Time
}

type VehicleExit interface {
	TicketID() string
	ExitTime() time.Time
}

type VehicleParkingDetails struct {
	vehicleType common.VehicleType
	entryTime   time.Time
	exitTime    time.Time
	ticketID    string
}

func (v *VehicleParkingDetails) EntryTime() time.Time {
	return v.entryTime
}

func (v *VehicleParkingDetails) VehicleType() common.VehicleType {
	return v.vehicleType
}

func (v *VehicleParkingDetails) ExitTime() time.Time {
	return v.exitTime
}

func (v *VehicleParkingDetails) TicketID() string {
	return v.ticketID
}

func NewVehicleEntry(vehicleType common.VehicleType) *VehicleParkingDetails {
	return &VehicleParkingDetails{
		entryTime:   time.Now(),
		vehicleType: vehicleType,
	}
}

func MockVehicleEntry(vehicleType common.VehicleType, entryTimeStr string, timeZone *time.Location) *VehicleParkingDetails {
	parsedEntryTime, err := time.ParseInLocation(common.DefaultTimeFormat, entryTimeStr, timeZone)
	if err != nil {
		panic(err)
	}
	return &VehicleParkingDetails{
		entryTime:   parsedEntryTime,
		vehicleType: vehicleType,
	}
}

func MockVehicleExit(ticketID string, exitTimeStr string, timeZone *time.Location) *VehicleParkingDetails {
	parsedExitTime, err := time.ParseInLocation(common.DefaultTimeFormat, exitTimeStr, timeZone)
	if err != nil {
		panic(err)
	}
	return &VehicleParkingDetails{
		exitTime: parsedExitTime,
		ticketID: ticketID,
	}
}

func NewVehicleExit(ticketID string) *VehicleParkingDetails {
	timeNow := time.Now()
	return &VehicleParkingDetails{
		exitTime: timeNow,
		ticketID: ticketID,
	}
}

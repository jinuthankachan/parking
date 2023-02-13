package main

import (
	"log"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/pkg"
	"github.com/jinut2/parking/services"
)

var parkingLot *services.ParkingLotService

func init() {
	var err error
	parkingLot, err = services.NewParkingLot(common.DefaultTimezone, 3, 2, 1, pkg.MallFeeCalculator())
	if err != nil {
		panic(err)
	}
}

func main() {
	ticket01, err := parkingLot.ParkVehicle(common.TwoWheeler)
	if err != nil {
		log.Fatalf("No entry possible due to %s", err)
	}
	log.Printf("Ticket Issued: \n %s", ticket01.Print())

	time.Sleep(time.Second)

	// Unpark the two wheeler
	receipt01, err := parkingLot.UnparkVehicle(ticket01.TicketID)
	if err != nil {
		log.Fatalf("Unable to mark exit due to %s", err)
	}
	log.Printf("Receipt Issued: \n %s", receipt01.Print())
}

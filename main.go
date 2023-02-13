package main

import (
	"log"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/services"
)

func main() {
	mallParkingLot, err := services.NewParkingLot("mall", common.DefaultTimezone, 3, 2, 1)
	if err != nil {
		panic(err)
	}

	ticket01, err := mallParkingLot.ParkVehicle(common.TwoWheeler)
	if err != nil {
		log.Fatalf("No entry possible due to %s", err)
	}
	log.Printf("Ticket Issued: \n %s", ticket01.Print())

	time.Sleep(time.Second)

	// Unpark the two wheeler
	receipt01, err := mallParkingLot.UnparkVehicle(ticket01.TicketID)
	if err != nil {
		log.Fatalf("Unable to mark exit due to %s", err)
	}
	log.Printf("Receipt Issued: \n %s", receipt01.Print())
}

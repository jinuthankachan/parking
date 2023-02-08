package main

import (
	"log"
	"time"

	"github.com/jinut2/parking/common"
	"github.com/jinut2/parking/models"
	"github.com/jinut2/parking/pkg"
)

func main() {
	timeZone, err := time.LoadLocation(common.DefaultTimezone)
	if err != nil {
		panic(err)
	}
	mallParkingLot := models.NewParkingLot(common.Mall, 3, 2, 1)
	mallParkingTickets := models.NewTickets()
	mallParkingReceipts := models.NewReceipts()
	mallParkingFeesModel := models.NewParkingLotFeesModel(common.Mall)

	// Park a two wheeler
	ticket01, err := pkg.ParkVehicle(mallParkingLot, mallParkingTickets, models.NewVehicleEntry(common.TwoWheeler), timeZone)
	if err != nil {
		log.Fatalf("No entry possible due to %s", err)
	}
	log.Printf("Ticket Issued: \n %+v", ticket01)

	time.Sleep(time.Second)

	// Unpark the two wheeler
	receipt01, err := pkg.UnparkVehicle(
		models.NewVehicleExit(ticket01.TicketNumber),
		mallParkingLot,
		mallParkingTickets,
		mallParkingFeesModel,
		mallParkingReceipts,
		timeZone,
	)
	if err != nil {
		log.Fatalf("Unable to mark exit due to %s", err)
	}
	log.Printf("Receipt Issued: \n %+v", receipt01)
}

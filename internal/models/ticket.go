package models

import "time"

type TicketModel struct {
	SpotID    string
	EntryTime time.Time
}

type TicketCounter interface {
	AllotTicket(spotID string, entryTime time.Time) (ticketID string, err error)
	TicketDetails(ticketID string) (*TicketModel, error)
}

type Tickets map[string]TicketModel

func NewTickets() *Tickets {
	Tickets := Tickets(make(map[string]TicketModel))
	return &Tickets
}

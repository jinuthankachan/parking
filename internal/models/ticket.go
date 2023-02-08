package models

import (
	"fmt"
	"sync"
	"time"
)

type TicketModel struct {
	SpotID    string
	EntryTime time.Time
}

type TicketCounter interface {
	AllotTicket(spotID string, entryTime time.Time) (ticketID string, err error)
	TicketDetails(ticketID string) (*TicketModel, error)
}

type Tickets struct {
	store map[string]TicketModel
	mu    sync.Mutex
}

func NewTickets() *Tickets {
	store := make(map[string]TicketModel)
	return &Tickets{
		store: store,
	}
}

func (t *Tickets) AllotTicket(spotID string, entryTime time.Time) (ticketID string, err error) {
	ticket := TicketModel{
		SpotID:    spotID,
		EntryTime: entryTime,
	}
	t.mu.Lock()
	ticketID = fmt.Sprintf("%d", len(t.store)+1)
	t.store[ticketID] = ticket
	t.mu.Unlock()
	return
}

func (t *Tickets) TicketDetails(ticketID string) (*TicketModel, error) {
	if ticket, ok := t.store[ticketID]; ok {
		return &ticket, nil
	}
	return nil, fmt.Errorf("error: ticket %s not found", ticketID)
}

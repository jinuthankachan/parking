package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinut2/parking/common"
)

type ReceiptModel struct {
	ExitTime time.Time
	Fees     *common.Currency
}

type ReceiptBook interface {
	GenerateReceipt(ticket *TicketModel, fees *common.Currency, exitTime time.Time) (receiptID string, err error)
}

type Receipts struct {
	store map[string]ReceiptModel
	mu    sync.Mutex
}

func NewReceiptBook() *Receipts {
	receiptsStore := make(map[string]ReceiptModel)
	return &Receipts{
		store: receiptsStore,
	}
}

func (r *Receipts) GenerateReceipt(ticket *TicketModel, fees *common.Currency, exitTime time.Time) (receiptID string, err error) {
	receipt := ReceiptModel{
		ExitTime: exitTime,
		Fees:     fees,
	}
	r.mu.Lock()
	receiptID = fmt.Sprintf("R-%d", len(r.store)+1)
	r.store[receiptID] = receipt
	r.mu.Unlock()
	return receiptID, nil
}

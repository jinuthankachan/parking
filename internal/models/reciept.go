package models

import (
	"time"

	"github.com/jinut2/parking/common"
)

type ReceiptModel struct {
	ExitTime time.Time
	Fees     common.Currency
}

type ReceiptGenerator interface {
	GenerateReceipt(ticket *TicketModel, fees *common.Currency) (receiptID string, err error)
}

type Receipts map[string]ReceiptModel

func NewReceipts() *Receipts {
	receipts := Receipts(make(map[string]ReceiptModel))
	return &receipts
}

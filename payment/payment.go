package payment

import (
	"fmt"

	"github.com/google/uuid"
)

type Order struct {
	ID     string
	Amount int
	Status string
}

func NewOrder(amount int) *Order {
	return &Order{ID: uuid.NewString(), Amount: amount, Status: "created"}
}

func (o Order) IsValid() bool {
	return o.Amount > 0
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) ProcessPayment(order *Order) error {
	if !order.IsValid() {
		return fmt.Errorf("invalid order")
	}
	if order.Status == "paid" {
		return fmt.Errorf("order is already paid")
	}
	order.Status = "paid"
	return nil
}

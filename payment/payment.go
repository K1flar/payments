package payment

import "errors"

type Order struct {
	Amount int
	Status string
}

func NewOrder(amount int) *Order {
	return &Order{Amount: amount, Status: "created"}
}

func (o *Order) IsValid() bool {
	return o.Amount >= 0
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) ProcessPayment(order *Order) error {
	if order.Status == "paid" {
		return errors.New("order is already paid")
	}

	if !order.IsValid() {
		order.Status = "payment_failed"
		return errors.New("error")
	}

	if order.Amount == 0 {
		order.Status = "confirmed"
		return nil
	}

	order.Status = "paid"
	return nil
}

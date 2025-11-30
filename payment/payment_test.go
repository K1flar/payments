package payment

import "testing"

func TestNewOrder(t *testing.T) {
	o := NewOrder(100)
	if o.Amount != 100 {
		t.Errorf("expected 100, got %d", o.Amount)
	}
}

func TestOrderInvalidAmount(t *testing.T) {
	o := NewOrder(-10)
	if o.IsValid() {
		t.Errorf("order with negative amount should be invalid")
	}
}

func TestOrderValidAmount(t *testing.T) {
	o := NewOrder(100)
	if !o.IsValid() {
		t.Errorf("order with positive amount should be valid")
	}
}

func TestNewPaymentService(t *testing.T) {
	ps := NewPaymentService()
	if ps == nil {
		t.Errorf("payment service should not be nil")
	}
}

func TestProcessPaymentValidOrder(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(100)

	err := ps.ProcessPayment(o)
	if err != nil {
		t.Errorf("expected no error for valid order, got %v", err)
	}
}

func TestProcessPaymentInvalidOrder(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(-10)

	err := ps.ProcessPayment(o)
	if err == nil {
		t.Errorf("expected error for invalid order")
	}
}

func TestOrderStatus(t *testing.T) {
	o := NewOrder(100)
	if o.Status != "created" {
		t.Errorf("expected status 'created', got %s", o.Status)
	}
}

func TestProcessPaymentChangesStatus(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(100)

	err := ps.ProcessPayment(o)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if o.Status != "paid" {
		t.Errorf("expected status 'paid', got %s", o.Status)
	}
}

func TestOrderHasID(t *testing.T) {
	o := NewOrder(100)
	if o.ID == "" {
		t.Errorf("order should have ID")
	}
}

func TestOrderUniqueIDs(t *testing.T) {
	o1 := NewOrder(100)
	o2 := NewOrder(200)

	if o1.ID == o2.ID {
		t.Errorf("order IDs should be unique")
	}
}

func TestProcessPaymentZeroAmount(t *testing.T) {
	ps := NewPaymentService()
	o := NewOrder(0)

	err := ps.ProcessPayment(o)
	if err == nil {
		t.Errorf("expected error for zero amount order")
	}
}

func TestProcessPaymentAlreadyPaidOrder(t *testing.T) {
	ps := NewPaymentService()
	order := NewOrder(200)
	if !order.IsValid() {
		t.Fatalf("order must be is valid")
	}

	// Первая оплата - должна пройти успешно
	err := ps.ProcessPayment(order)
	if err != nil {
		t.Fatalf("unexpected error processing first payment: %v", err)
	}

	// Вторая оплата - должна вернуть ошибку
	err = ps.ProcessPayment(order)
	if err == nil {
		t.Errorf("expected error when processing payment for already paid order")
	}

	// Проверяем, что статус остался "paid"
	if order.Status != "paid" {
		t.Errorf("expected status to remain 'paid', got %s", order.Status)
	}
}

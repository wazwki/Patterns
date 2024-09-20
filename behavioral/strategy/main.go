package main

import "fmt"

// PaymentStrategy - интерфейс стратегии оплаты
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCard - стратегия оплаты кредитной картой
type CreditCard struct {
	name   string
	number string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card (%s)\n", amount, c.name)
}

// PayPal - стратегия оплаты через PayPal
type PayPal struct {
	email string
}

func (p *PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal (%s)\n", amount, p.email)
}

// Cash - стратегия оплаты наличными
type Cash struct{}

func (c *Cash) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Cash\n", amount)
}

// Order - контекст, который использует стратегии оплаты
type Order struct {
	amount  float64
	payment PaymentStrategy
}

func (o *Order) SetPaymentStrategy(payment PaymentStrategy) {
	o.payment = payment
}

func (o *Order) Pay() {
	o.payment.Pay(o.amount)
}

func main() {
	// Создаем заказ на сумму 50
	order := &Order{amount: 50}

	// Оплата через кредитную карту
	order.SetPaymentStrategy(&CreditCard{name: "Alice", number: "1234-5678-9876-5432"})
	order.Pay()

	// Оплата через PayPal
	order.SetPaymentStrategy(&PayPal{email: "alice@example.com"})
	order.Pay()

	// Оплата наличными
	order.SetPaymentStrategy(&Cash{})
	order.Pay()
}

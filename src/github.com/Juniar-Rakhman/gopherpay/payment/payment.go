package payment

type PaymentOption interface {
	ProcessPayment(amount float32) bool
}

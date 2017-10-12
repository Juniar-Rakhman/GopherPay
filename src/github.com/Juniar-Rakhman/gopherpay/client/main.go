package main

import (
	"github.com/Juniar-Rakhman/gopherpay/payment"
)

func main() {
	var option payment.PaymentOption

	account := payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(account.OwnerName(), 500)
}

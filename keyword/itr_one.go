package keyword

import (
	"fmt"

	"github.com/stretchr/testify/require"
)

type PaymentMethod interface {
	Pay(amount int)
}

func BulkPay(bulkPay []PaymentMethod, payOut []int) {
	require.Equal(nil, len(bulkPay), len(payOut), "bulkPay and payOut must be of same length")

	for index, value := range payOut {
		bulkPay[index].Pay(value)
	}
}

type CardDetails struct {
	HolderName string
	CardNumber int
	Balance    int
	Type       string
}

func (c *CardDetails) Pay(amount int) {
	c.Balance -= amount
	fmt.Printf(
		"User Name: %s, Card Number: %v, Remaining Balance: %d, Card Type: %s\n",
		c.HolderName, c.CardNumber, c.Balance, c.Type,
	)
}

func ItrOne() {
	creditInfo := &CardDetails{"Aravindh", 14839459384935, 10000, "Credit Card"}
	debitInfo := &CardDetails{"Aravindh", 14839459384935, 20000, "Debit Card"}
	gpayInfo := &CardDetails{"Aravindh", 14839459384935, 20000, "G Pay"}
	paypalInfo := &CardDetails{"Aravindh", 14839459384935, 20000, "PayPal"}

	paymentMethods := []PaymentMethod{creditInfo, debitInfo, gpayInfo, paypalInfo}
	payments := []int{10, 20, 30}

	BulkPay(paymentMethods, payments)


	fmt.Println("Interface")
}

package card

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret string
	Key string
	Currency string
}

type Transaction struct {
	TransactionID string
	Amount int
	Currency string
	LastFour string
	BankReturnCode string
}

func (c *Card) CreatePaymentIntent(currency string , amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code) 
		}
		return nil, msg, err
	}
	return pi, "", nil
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount) 
}

func cardErrorMessage(code stripe.ErrorCode) string {
	msg := ""

	switch code {
		case stripe.ErrorCodeCardDeclined:
			msg = "Your card was declined. Please check your card details and try again."
		case stripe.ErrorCodeExpiredCard:
			msg = "Your card has expired. Please use a different card."
		default: 
			msg = "An error occurred while processing your card. Please try again later."
	}

	return msg
}
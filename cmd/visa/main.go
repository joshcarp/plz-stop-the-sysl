package main

import (
	"context"
	"fmt"
	"log"

	visa "github.com/anz-bank/conf-demo/gen/pkg/servers/Visa"
)

// AppConfig ...
type AppConfig struct{}

func main() {
	log.Fatal(visa.Serve(
		context.Background(),
		func(ctx context.Context, config AppConfig) (*visa.ServiceInterface, error) {
			return &visa.ServiceInterface{
				PostPay: func(ctx context.Context, req *visa.PostPayRequest, client visa.PostPayClient) (*visa.ResponseData, error) {
					return &visa.ResponseData{
						Response: fmt.Sprintf(
							"Account %s successfully did a payment in the amount on %d %s",
							req.Request.CreditCardNumber,
							req.Request.AmountInSubUnits,
							req.Request.Currency,
						),
					}, nil
				},
			}, nil
		},
	))
}

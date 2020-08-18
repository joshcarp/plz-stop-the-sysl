package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	paymentserver "github.com/anz-bank/conf-demo/gen/pkg/servers/PaymentServer"
	"github.com/anz-bank/conf-demo/gen/pkg/servers/PaymentServer/mastercard"
	visa "github.com/anz-bank/conf-demo/gen/pkg/servers/PaymentServer/visa"
)

// AppConfig ...
type AppConfig struct {
	StartingBalance int64 `yaml:"startingBalance"`
}

func main() {
	log.Fatal(paymentserver.Serve(context.Background(),
		func(ctx context.Context, config AppConfig) (*paymentserver.ServiceInterface, error) {
			return &paymentserver.ServiceInterface{
				GetBalanceList: func(
					ctx context.Context,
					req *paymentserver.GetBalanceListRequest,
					client paymentserver.GetBalanceListClient,
				) (*paymentserver.Balance, error) {
					var amt int64
					if err := client.GetBalance.QueryRowContext(ctx).Scan(&amt); err != nil {
						if err == sql.ErrNoRows {
							amt = config.StartingBalance
							if _, err = client.InsertBalance.ExecContext(ctx, 1, amt); err == nil {
								goto ok
							}
						}
						return nil, err
					}
				ok:
					return &paymentserver.Balance{Amount: amt}, nil
				},

				PostPay: func(
					ctx context.Context,
					req *paymentserver.PostPayRequest,
					client paymentserver.PostPayClient,
				) (*paymentserver.PaymentResponse, error) {
					resp := &paymentserver.PaymentResponse{}
					switch {
					case strings.HasPrefix(req.Request.AccountNumber, "4"):
						r, err := client.VisaPostPay(ctx, &visa.PostPayRequest{
							Request: visa.PaymentData{
								AmountInSubUnits: int64(req.Request.Amount),
								CreditCardNumber: req.Request.AccountNumber,
								Currency:         "AUD",
							},
						})
						if err != nil {
							return resp, err
						}
						resp.Message = r.Response
						return resp, nil
					case strings.HasPrefix(req.Request.AccountNumber, "5"),
						strings.HasPrefix(req.Request.AccountNumber, "2"):
						state := mastercard.AustralianState("vic")
						_, err := client.MastercardPostPay(ctx, &mastercard.PostPayRequest{
							Request: mastercard.Payment{
								AmountInSubUnits: int64(req.Request.Amount),
								CreditCardNumber: req.Request.AccountNumber,
								Currency:         "AUD",
								State:            &state,
							},
						})
						if err != nil {
							return resp, err
						}
					default:
						return resp, fmt.Errorf("Account number is invalid: %s", req.Request.AccountNumber)
					}
					_, err := client.MakePayment.ExecContext(ctx, 1)
					return resp, err
				},
			}, nil
		},
	))
}

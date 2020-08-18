// Code generated by sysl DO NOT EDIT.
package paymentserver

import (
	"context"

	"github.com/anz-bank/conf-demo/gen/pkg/servers/PaymentServer/mastercard"
	"github.com/anz-bank/conf-demo/gen/pkg/servers/PaymentServer/visa"
	"github.com/anz-bank/sysl-go/catalogservice"
	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"

	"github.com/go-chi/chi"
)

// DownstreamClients for PaymentServer
type DownstreamClients struct {
	mastercardClient *mastercard.Client
	visaClient       *visa.Client
}

// BuildRestHandlerInitialiser ...
func BuildRestHandlerInitialiser(
	serviceInterface ServiceInterface,
	callback core.RestGenCallback,
	downstream *DownstreamClients,
) (handlerinitialiser.HandlerInitialiser, error) {
	serviceHandler, err := NewServiceHandler(callback, &serviceInterface, downstream.mastercardClient, downstream.visaClient)
	if err != nil {
		return nil, err
	}
	return NewServiceRouter(callback, serviceHandler), nil
}

// BuildDownstreamClients ...
func BuildDownstreamClients(cfg *config.DefaultConfig) (*DownstreamClients, error) {
	var err error = nil
	mastercardHTTPClient, mastercardErr := core.BuildDownstreamHTTPClient("mastercard", &cfg.GenCode.Downstream.(*DownstreamConfig).Mastercard)
	visaHTTPClient, visaErr := core.BuildDownstreamHTTPClient("visa", &cfg.GenCode.Downstream.(*DownstreamConfig).Visa)
	if mastercardErr != nil {
		return nil, mastercardErr
	}

	if visaErr != nil {
		return nil, visaErr
	}

	mastercardClient := mastercard.NewClient(mastercardHTTPClient, cfg.GenCode.Downstream.(*DownstreamConfig).Mastercard.ServiceURL)
	visaClient := visa.NewClient(visaHTTPClient, cfg.GenCode.Downstream.(*DownstreamConfig).Visa.ServiceURL)

	return &DownstreamClients{mastercardClient: mastercardClient,
		visaClient: visaClient,
	}, err
}

// Serve starts the server.
//
// createService must be a function with the following signature:
//
//   func(ctx context.Context, config AppConfig) (*paymentserver.ServiceInterface, error)
//
// where AppConfig is a type defined by the application programmer to
// hold application-level configuration.
func Serve(
	ctx context.Context,
	createService interface{},
) error {
	return core.Serve(
		ctx,
		&DownstreamConfig{}, createService, &ServiceInterface{},
		func(cfg *config.DefaultConfig, serviceIntf interface{}) (chi.Router, error) {
			serviceInterface := serviceIntf.(*ServiceInterface)

			genCallbacks := common.DefaultCallback()

			clients, err := BuildDownstreamClients(cfg)
			if err != nil {
				return nil, err
			}

			serviceHandler, err := NewServiceHandler(
				genCallbacks,
				serviceInterface,
				clients.mastercardClient,
				clients.visaClient,
			)
			if err != nil {
				return nil, err
			}

			// Service Router
			router := chi.NewRouter()
			serviceRouter := NewServiceRouter(genCallbacks, serviceHandler)
			serviceRouter.WireRoutes(ctx, router)
			catalogservice.Enable(serviceRouter, router, AppSpec)
			return router, nil
		},
	)
}

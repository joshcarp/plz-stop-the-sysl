# conf-demo

Unconf demo for sysl

## Prerequisites

- [Sysl v0.11.0 or later ](https://sysl.io/docs/install/)
- go 1.13
- Docker

## File structure 

`api/`: contains all API specifications for the generated application

`gen/`: contains all the generated code for the service

`cmd/sizzle`: runs the generated server

## running code generation
- in [Makefile](Makefile) specify the following:
```
SYSLFILE = api/sizzle.sysl
APPS = PaymentServer
```
- run `make` to regenerate application code


## Run the service:

`go run ./cmd/sizzle`
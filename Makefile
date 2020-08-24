SYSLFILE = api/plz.sysl
APPS = myserver

-include local.mk
include extra/codegen.mk


input=/dep/api
output=/dep/plzserver
proto:
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I.$(input) --go_out=paths=source_relative:$(output) api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I.$(input) --go-grpc_out=paths=source_relative:$(output) api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I.$(input) --go-grpc_out=paths=source_relative:$(output) api.proto
	docker run --rm -v $$(pwd):/dep:rw anzbank/protoc-gen-sysl:v0.0.20 -I.$(input) --sysl_out=$(output) api.proto


clean:
	rm -rf gen/pkg/servers/myserver

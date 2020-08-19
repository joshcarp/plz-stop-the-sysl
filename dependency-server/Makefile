

proto:
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc $(INCLUDE) -I./dep/proto/ --go_out=paths=source_relative:/dep/proto/ api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc $(INCLUDE) -I./dep/proto/ --go-grpc_out=paths=source_relative:/dep/proto/ api.proto
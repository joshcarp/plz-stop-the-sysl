
.PHONY:proto
proto:
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/mypkg/ --go_out=paths=source_relative:/dep/mypkg/ api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/mypkg/ --go-grpc_out=paths=source_relative:/dep/mypkg/ api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/mypkg/ --go-grpc_out=paths=source_relative:/dep/mypkg/ api.proto
	docker run --rm -v $$(pwd):/dep:rw anzbank/protoc-gen-sysl:v0.0.20 -I./dep/mypkg/ --sysl_out=/dep/mypkg/ api.proto
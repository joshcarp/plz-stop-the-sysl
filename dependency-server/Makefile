
.PHONY:proto
proto:
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/depserver/ --go_out=paths=source_relative:/dep/depserver/ api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/depserver/ --go-grpc_out=paths=source_relative:/dep/depserver/ api.proto
	docker run --rm -v $$(pwd):/dep:rw joshcarp/protoc -I./dep/depserver/ --go-grpc_out=paths=source_relative:/dep/depserver/ api.proto
	docker run --rm -v $$(pwd):/dep:rw anzbank/protoc-gen-sysl:v0.0.20 -I./dep/depserver/ --sysl_out=/dep/depserver/ api.proto
	mv depserver/index.sysl ../api/downstream.sysl
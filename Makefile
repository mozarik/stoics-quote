.PHONY: proto-gen
proto-gen:
	@echo "Generating proto files"; \
	protoc --go_out=plugins=grpc:. *.proto;

protoc --go_out=. \
    --go-grpc_out=.  \
    proto/api-gateway.proto
	


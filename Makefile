.PHONY: proto-gen
proto-gen:
	@echo "Generating proto files"; \
	protoc --go_out=plugins=grpc:. *.proto;

.PHONY: run-main-service
run-main-service:
	@echo "Running main service"; \
	go run main-svc/main.go;
	
.PHONY: run-auth-api-gateway-svc
run-auth-api-gateway-svc:
	@echo "Running auth api gateway service"; \
	go run auth-api-gateway-svc/main.go;


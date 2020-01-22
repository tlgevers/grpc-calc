stubs: ## Generate stubs
	protoc -I proto proto/calc.proto --go_out=plugins=grpc:pkg/calc

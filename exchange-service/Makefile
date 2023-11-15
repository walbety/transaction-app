PROTO_SRC_PATH := ./tools/protos

run:
	go run $$(find cmd/ -name main.go) --config-path=configs/config.json

test:
	go test -tags='test' ./...

protobuf-exchange:
	protoc --go_out=. $(PROTO_SRC_PATH)/exchange.proto
	protoc --go-grpc_out=. $(PROTO_SRC_PATH)/exchange.proto

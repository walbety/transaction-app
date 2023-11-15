PROTO_SRC_PATH := ./tools/protos

run:
	go run $$(find cmd/ -name main.go) --config-path=configs/config.json

test:
	go test -tags='test' ./...

protobuf-exchange:
	protoc --go_out=. $(PROTO_SRC_PATH)/exchange.proto
	protoc --go-grpc_out=. $(PROTO_SRC_PATH)/exchange.proto

cover-tree:
	go install github.com/nikolaydubina/go-cover-treemap@latest
	go test -tags='test' -coverprofile cover.out ./...
	go-cover-treemap -coverprofile cover.out > cover.svg
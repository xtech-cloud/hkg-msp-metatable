.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/shared.proto

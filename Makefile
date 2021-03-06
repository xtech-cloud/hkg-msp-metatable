.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/vocabulary.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/source.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/schema.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/metatable/format.proto

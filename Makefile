filename ?= core_engine.proto
folder ?= core-engine

generate-rpc:
	protoc -I protobufs --go_out=grpc/${folder} --go_opt=paths=source_relative --go-grpc_out=grpc/${folder} --go-grpc_opt=paths=source_relative protobufs/${filename}
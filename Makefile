generate-server:
	@protoc --go-grpc_out=pkg/protos --go_out=pkg/protos pkg/protos/coins.proto

tidy:
	go mod tidy
create_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=protos --go-grpc_opt=paths=source_relative \
    helloworld.proto
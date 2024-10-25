generate:
	protoc --go_out=./proto/protofiles --go_opt=paths=source_relative \
		--go-grpc_out=./proto/protofiles --go-grpc_opt=paths=source_relative \
		--proto_path=./proto \
    	./proto/*.proto
gen:
	protoc --proto_path=proto proto/*.proto \
	--go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import

	protoc --proto_path=proto proto/edges/*.proto \
	--go_out=. --go_opt=paths=import \
	--go-grpc_out=. --go-grpc_opt=paths=import

	protoc --proto_path=proto proto/slots/*.proto \
	--go_out=. --go_opt=paths=import \
	--go-grpc_out=. --go-grpc_opt=paths=import

	rm -r pb || true
	cp -r github.com/snple/kirara/pb .
	rm -r github.com

.PHONY: gen

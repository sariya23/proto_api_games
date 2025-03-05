gen_auth:
	protoc -I proto proto/auth/auth.proto \
	--go_out=./gen/auth \
	--go_opt=paths=source_relative \
	--go-grpc_out=./gen/auth \
	--go-grpc_opt=paths=source_relative 
.PHONY: gen_auth

clean:
	rm -rf gen/*
gen_auth:
	mkdir -pp gen/auth && \
	protoc -I ./proto proto/auth/auth.proto \
	--go_out ./gen \
	--go_opt paths=source_relative \
	--go-grpc_opt paths=source_relative \
	--go-grpc_out ./gen \
	--grpc-gateway_out ./gen \
	--grpc-gateway_opt paths=source_relative
.PHONY: gen_auth

clean:
	rm -rf gen/*
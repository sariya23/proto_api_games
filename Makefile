gen_auth:
	mkdir -pp gen/auth && \
	protoc -I ./proto proto/auth/auth.proto \
	--go_out ./gen \
	--go_opt paths=source_relative \
	--go-grpc_opt paths=source_relative \
	--go-grpc_out ./gen \
	--grpc-gateway_out ./gen \
	--swagger_out=:swagger \
	--grpc-gateway_opt paths=source_relative
.PHONY: gen_auth

gen_game:
	mkdir -pp gen/game && \
	protoc -I ./proto proto/game/game.proto \
	--go_out ./gen \
	--go_opt paths=source_relative \
	--go-grpc_opt paths=source_relative \
	--go-grpc_out ./gen \
	--grpc-gateway_out ./gen \
	--openapiv2_out ./swagger
	--grpc-gateway_opt paths=source_relative

clean:
	rm -rf gen/*
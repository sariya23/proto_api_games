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

gen_gamev2:
	mkdir -p gen/game && \
	protoc -I ./proto proto/gamev2/gamev2.proto \
		--go_out ./gen \
		--go_opt paths=source_relative \
		--go-grpc_out ./gen \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./swagger

clean:
	rm -rf gen/*
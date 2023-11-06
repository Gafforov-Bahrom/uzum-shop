DATABASE_DSN=postgres://postgres:root@localhost:5432/uzum_shop?sslmode=disable

generate:
	rm -rf ./pkg/grpc && mkdir ./pkg/grpc
	protoc --go_out=./pkg/grpc --go-grpc_out=./pkg/grpc api/*.proto

create-migration:
	migrate create -ext=sql -dir=migrations migration_name

migrate-up:
	migrate -source file://migrations -database "${DATABASE_DSN}" up

migrate-down:
	migrate -source file://migrations -database "${DATABASE_DSN}" down



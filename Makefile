.PHONY: postgres createdb dropdb psql migrateup migrateup1 migratedown migratedown1 sqlc mock test server

postgres:
	docker run --name postgres16.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine3.19

createdb:
	docker exec -it postgres16.2 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16.2 dropdb simple_bank

psql:
	docker exec -it postgres16.2 psql -U root -d simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

mock:
	mockgen -source db/sqlc/store.go -destination db/mock/store.go

test:
	go test -v -cover ./...

server:
	go run main.go

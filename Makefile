.PHONY: postgres createdb dropdb psql migrateup migratedown sqlc

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

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

postgres:
	docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root e_commerce

dropdb:
	docker exec -it postgres16 dropdb e_commerce

migrateup:
	migrate -database "postgresql://root:root@localhost:5432/e_commerce?sslmode=disable" -path db/migrations -verbose up

migrateup1:
	migrate -database "postgresql://root:root@localhost:5432/e_commerce?sslmode=disable" -path db/migrations -verbose up 1

migratedown:
	migrate -database "postgresql://root:root@localhost:5432/e_commerce?sslmode=disable" -path db/migrations -verbose down

migratedown1:
	migrate -database "postgresql://root:root@localhost:5432/e_commerce?sslmode=disable" -path db/migrations -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1
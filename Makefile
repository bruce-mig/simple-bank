postgres: 
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sudo /home/migeri/.gvm/pkgsets/go1.21.4/global/bin/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	sudo /home/migeri/.gvm/pkgsets/go1.21.4/global/bin/mockgen -package mockdb -destination db/mock/store.go github.com/bruce-mig/simple-bank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1
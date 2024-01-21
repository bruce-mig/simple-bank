DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres: 
	docker run --name bank-postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sudo /home/migeri/.gvm/pkgsets/go1.21.4/global/bin/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock :
	mockgen -package mockdb -destination db/mock/store.go github.com/bruce-mig/simple-bank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/bruce-mig/simple-bank/worker TaskDistributor

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

redis:
	docker run --name redis -p 6379:6379 -d redis:7.2-alpine


.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1 new_migration proto db_docs db_schema redis
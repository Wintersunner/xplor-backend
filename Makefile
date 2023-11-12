include app.env

migrate:
	migrate -path db/migration \
  -database "${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
  -verbose up

rollback:
	migrate -path db/migration \
  -database "${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
  -verbose down 1

resetdb:
	migrate -path db/migration \
  -database "${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
  -verbose down --all

refresh: rollback migrate

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Wintersunner/xplor/db/sqlc Store

build:
	go build -o main main.go

run:
	go run main.go
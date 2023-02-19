postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root music_player

dropdb:
	docker exec -it postgres14 dropdb music_player

migrateup:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5432/music_player?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5432/music_player?sslmode=disable" -verbose down

sqlcinit:
	docker run --rm -v "C:\Users\chien\Projects\MusicPlayerBE:/src" -w /src kjconroy/sqlc:1.16.0 init

sqlc:
	docker run --rm -v "C:\Users\chien\Projects\MusicPlayerBE:/src" -w /src kjconroy/sqlc:1.16.0 generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/Chien179/SimpleBank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
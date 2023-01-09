
DB=postgresql://postgres:secret@localhost:5432/aozora-rodoku?sslmode=disable

db-apply:
	atlas schema apply -u ${DB} --to file://db/schema.hcl

db-inspect:
	atlas schema inspect -u ${DB}

# コンテナ起動
up:
	cd docker && docker-compose up
# DBを作成
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simplebank

# DB削除
dropdb:
	docker exec -it postgres dropdb -Upostgres simplebank

# マイグレーション
migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

# スキーマーの作成
createsc:
	migrate create -ext sql -dir db/migration -seq ${name}

.PHONY: up createdb dropdb migrateup migratedown createsc


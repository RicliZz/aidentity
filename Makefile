include .env
export

build:
	@go build -o ./bin/app ./cmd/main.go

run:build
	@./bin/app

new_migrate:
	@migrate create -ext sql -dir db/migrations -seq ${name}

migrate_up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down

collect_quality:
	@go run /home/ricliz/GolandProjects/sobes/main.go
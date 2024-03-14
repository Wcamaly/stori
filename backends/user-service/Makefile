DB_USER := $(or $(DB_USER),root)
DB_PASSWORD := $(or $(DB_PASSWORD),rootPassword)
DB_NAME := $(or $(DB_NAME),stori_service)
DB_HOST := $(or $(DB_HOST),localhost)
DB_PORT := $(or $(DB_PORT),5432)

createdb:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

migrate:
	export PGPASSWORD=$(DB_PASSWORD); \
	for file in database/V*; do \
		psql -U $(DB_USER)  -h $(DB_HOST) -p $(DB_PORT) -d $(DB_NAME) -a -f $$file; \
	done

tidy:
	go mod tidy
	go mod vendor

run:
	go run main.go


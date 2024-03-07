generate:
	go generate ./...

test:
	go test ./... -cover -v -coverprofile=coverage.out

show-out:
	go tool cover -html=.coverage.out

tidy:
	go mod tidy
	go mod vendor

run:
	go run main.go

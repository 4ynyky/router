build:
	go build cmd/radius/main.go

run:
	go build cmd/radius/main.go
	./main -f ./internal/config/config.json
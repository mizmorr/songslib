
run:
		cd cmd/bin; go run main.go

test:
	go test -v -cover ./...

swag:
	swag init -g cmd/bin/main.go -o docs

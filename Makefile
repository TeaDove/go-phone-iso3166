check:
	pre-commit run -a
	go mod tidy
	go test -v ./...

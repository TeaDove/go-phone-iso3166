test:
	go test -v ./...

check: test
	pre-commit run -a
	go mod tidy

build:
	@go build -o bin/resume-comp-backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/resume-comp-backend
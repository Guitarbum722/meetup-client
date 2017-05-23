build:
	go build -v ./...

test:
	go test -v ./...

generate-mocks:
	mockery -dir=./ -all

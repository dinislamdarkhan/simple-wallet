run:
	go run main.go

lint:
	golangci-lint run

test:
	go test ./...

fumpt:
	gofumpt -l -w .

mock:
	mockery --all --keeptree

remock:
	rm -rf ./mocks
	mockery --all --keeptree

swagger: check-swagger
	swagger generate spec -o ./swagger.json -m

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.json

check: lint test
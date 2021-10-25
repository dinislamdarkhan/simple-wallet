# Simple Wallet API

Simple Wallet API is project which is a server and allows the user to make transactions, get account balance and transaction history. Written with GO version 1.17
## Documentation

## Installation
Clone repository and run
```bash
go run main.go
```

## Usage
For full usage of project recommended install some  packages:
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [gofumpt](https://github.com/mvdan/gofumpt)
- [swagger](https://github.com/go-swagger/go-swagger)
- [mockery](https://github.com/vektra/mockery)

Check linters
```bash
golangci-lint run
```
Run tests
```bash
go test ./...
```
Format files
```bash
gofumpt -l -w .
```
Generate swagger
```bash
swagger generate spec -o ./swagger.json -m
```
Serve swagger
```bash
swagger serve -F=swagger swagger.json
```
Mock files
```bash
swagger serve -F=swagger swagger.json
```
Remock files
```bash
rm -rf ./mocks
mockery --all --keeptree
```
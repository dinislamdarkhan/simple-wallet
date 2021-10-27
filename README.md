# Simple Wallet API

Simple Wallet API is project which is a server and allows the user to make transactions, get account balance and transaction history. Written with GO version 1.17
## Documentation
All development process can be found in [Trello Table](https://trello.com/b/0Zzkpg9W/simple-wallet)
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

## Git flow
Production branch - ```master```

Development branch - ```qa```

Branch naming:
- ```feature/``` : only for new code
- ```hotfix/``` : only for fix
- ```release/``` : release qa code to master
- ```config/```: only for changing files without code

All development part works in feature branches. One feature = One business task.

```feature/``` branches can be started only from ```qa``` branch and PR-ted into ```qa```

Every feature branch can be decomposited for few subbranches, like - ```feature/go-transaction``` has ```go-transaction/readme``` subbranch for delegate tasks and simplify every Pull Request.

```hotfix/``` branches can be started from ```qa``` or ```master```, if started from master, we need make 2 PR, one to master, one to qa for actualization

```release/``` branches can be started only from ```qa```, after testing, bug fixing in envs we need make PR to ```master```
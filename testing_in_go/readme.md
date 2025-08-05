# Testing in go #

```bash
go version
```

## Extensions ##

- go, then ctrl+shift+p, select *Go: Install/Update Tools*, select all and hit ok.
- gotemplate-syntax

Also install docker...

## Test ##

```bash
go test .
go test -v .
go test -cover .
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

go test -run Test_isPrime
go test -v -run Test_isPrime
```

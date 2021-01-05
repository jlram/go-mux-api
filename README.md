# About

This is a personal project I started to improve my Go skills in a backend sense, including:

- [net/http](https://golang.org/pkg/net/http/)
- [gorilla/mux](https://github.com/gorilla/mux)
- [squirrel](https://github.com/Masterminds/squirrel) (pending)
- Test Driven Development (TDD)
- Continuous Integration and Continuous Development (CI/CD)

It firstly consists in a product CRUD application, but would like to do expand it in the future ([here](https://github.com/victorst79/food-scanner))

# Testing

### Run test instance of postgreSQL
`docker run -it -p 5432:5432 -e POSTGRES_PASSWORD=openpgpwd -d postgres`

### Run tests
`cd src && go test -v`

# Running

### Option 1: Run project

`go run .`

### Option 2: Alternatively, generate executable file

`go build`, then run /src/src.exe

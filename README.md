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

### Output should be like this:

```
Up and running!
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.00s)
=== RUN   TestNonExistentProduct
--- PASS: TestNonExistentProduct (0.00s)
=== RUN   TestCreateProduct
--- PASS: TestCreateProduct (0.00s)
=== RUN   TestRetrieveProduct
--- PASS: TestRetrieveProduct (0.00s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.00s)
--- FAIL: TestUpdateProduct (0.00s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.00s)
PASS
ok
```

# Running

### Option 1: Run project

`cd src && go run .`

### Option 2: Alternatively, generate executable file

`cd src && go build`, then run /src/src.exe

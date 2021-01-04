## Steps

### Run test instance of postgreSQL
`docker run -it -p 5432:5432 -d postgres`

### Run tests
`cd api && go test -v`
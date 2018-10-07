## Dependencies
1. `github.com/go-pg/pg`
2. `github.com/go-pg/migrations` 


## Instructions
1. You need to create database `gnippets` before running this example.
2. `cd ./migrations && go run *.go`
3. at the root lvl run `go run main.go`

## Migrate up/down/seed commands
```bash
> psql -c "CREATE DATABASE gnippets"
CREATE DATABASE

> go run *.go version
version is 0

> go run *.go
Creating cars table
Adding model to cars table
Seeding cars table
migrated from version 0 to 3

> go run *.go version
version is 3

> go run *.go reset
Truncating cars table
Removing model colmn from cars table
Dropping cars table
migrated from version 3 to 0

> go run *.go up 2
Creating cars table
Adding model to cars table
migrated from version 0 to 2

> go run *.go
Seeding cars table
migrated from version 2 to 3

> go run *.go down
Truncating cars table
migrated from version 3 to 2

> go run *.go version
version is 2

> go run *.go set_version 1
migrated from version 2 to 1

> go run *.go create add fuel type to cars
created migration 4_add_fuel_type_to_cars.go
```

## Transactions

If you'd want to wrap the whole run in a big transaction, which may be the case if you have multi-statement migrations, the code in `main.go` should be slightly modified:

```go
var oldVersion, newVersion int64

err := db.RunInTransaction(func(tx *pg.Tx) (err error) {
    oldVersion, newVersion, err = migrations.Run(tx, flag.Args()...)
    return
})
if err != nil {
    exitf(err.Error())
}
```


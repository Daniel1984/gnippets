### USAGE
`go run main.go` or `go build main.go && ./main` to se default setting of env
var in action.

`DB_CONN=postgres:conn go run main.go` to accept env variable.
You can also use `os.Unsetenv(key)` to unset env variable.

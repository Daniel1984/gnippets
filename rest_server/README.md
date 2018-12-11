## Dependencies
1. `github.com/go-pg/pg`
2. `github.com/julienschmidt/httprouter`

## Instructions
In real application you should add `config.json` to `.gitignore` for security reasons

## Actions

1. create user: `curl -d '{"name":"foo", "role":"gopher"}' -H "Content-Type: application/json" -X POST "http://localhost:8080/users"`
2. list all users: `curl localhost:8080/users`
3. get user by id: `curl localhost:8080/users/1`
4. patch user `curl -d '{"name":"carl shelby"}' -H "Content-Type: application/json" -X PATCH "http://localhost:8080/users/1"`
5. delete user `curl -X DELETE "http://localhost:8080/users/:userId`

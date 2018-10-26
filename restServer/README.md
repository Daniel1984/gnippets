## Dependencies
1. `github.com/go-pg/pg`
2. `github.com/julienschmidt/httprouter`

## Instructions
In real application you should add `config.json` to `.gitignore` for security reasons

##Actions

1. `POST` create user: `curl -d '{"name":"foo", "role":"gopher"}' -H
   "Content-Type: application/json" -X POST "http://localhost:8080/users"`
2. `GET` all users: `curl localhost:8080/users`
3. `GET` user by id: `curl localhost:8080/users/1`

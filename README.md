## Usage
### rename [.env.example](.env.example) to .env
install dependencies
``` bash
go mod tidy
```
start up the database
```bash
docker compose up -d
```
for the first time you load up the server, you need to migrate the database
```bash
go run ./cmd/api migrate
```
 seed database with stuff (admin user, exercises, collections, routines)
```bash
go run ./cmd/api seed
```
run the server
```bash
go run ./cmd/api
```
## <a href="./internal/routes/routes.go">Routes</a>

### Todos
[See older commits to view what has been done so far](https://github.com/xyztavo/go-gym/commits/main)

- [ ] seed also with arnold split and bros plit 


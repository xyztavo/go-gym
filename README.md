## Usage for Dev 
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
## Production
this project has a [Dockerfile](./Dockerfile), it only spins up the GO API, but you can use external database solutions (like the [neon serverless postgres](https://neon.tech) or even vercel db)
## <a href="./internal/routes/routes.go">Routes</a>

### Todos
[See older commits to view what has been done so far](https://github.com/xyztavo/go-gym/commits/main)

- [x] seed also with arnold split 
- [ ] seed also with and bro split 


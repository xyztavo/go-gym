### Todos
[...] [See older commits to view what has been done so far]('https://https://github.com/xyztavo/go-gym/commits/main/')
- [x] create exercises routines collection table, gym admins can set gym routines collection
- [x] users from gym can GET gym routines 
- [x] users can GET collections by routines ID  
- [x] users can GET exercises reps by collections ID 
- [x] users has a route that gets all info about the gym in one route (plans, routines...)
- [x] user route to get info about the user plan (active plan, lastpayment, when plan expires....)
- [x] give plans, routines, collections img attr in db and refactor everything(pain)
- [x] delete exercises-reps-collections by id
- [x] Delete routine collections
- [ ] clean up with exercises that really matter in seed
- [ ] seed with collections 
- [ ] seed with routines
- [ ] frontend (80% done)


## Usage
install dependencies
``` bash
go mod tidy
```
for the first time you load up the server, you need to migrate the database
```bash
go run ./cmd/api migrate
```
optional - seed database with stuff (exercises, collections, routines)
```bash
go run ./cmd/api seed
```
run the server
```bash
go run ./cmd/api
```

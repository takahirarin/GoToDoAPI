# GoToDoAPI

```bash
docker compose build
docker compose up -d
docker compose exec db bash
>psql -U rin -d rin
psql >CREATE DATABASE test_database;
psql >exit
docker compose exec todo_app sh
go run cmd/main.go -option=migrate
```

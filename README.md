# GoToDoAPI

初回立ち上げ
```bash
docker compose build
docker compose up -d
docker compose exec db bash
> psql -U rin -d rin
psql > CREATE DATABASE test_database;
psql > exit
docker compose exec todo_app sh
go run cmd/main.go -option=migrate
```
一覧取得
```bash
curl -i localhost:8080/fetch-todos
```

追加
```bash
curl -X POST "http://localhost:8080/add-todo" -d 'name=あああ&status=いいい'
```

変更
```bash
curl -X POST "http://localhost:8080/change-todo" -d 'id=ううう&status=作業中'
```

削除
```bash
curl -X POST "http://localhost:8080/delete-todo" -d 'id=えええ'
```

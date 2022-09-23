### Steps

go run -mod=mod entgo.io/ent/cmd/ent init Individual IncomeBracket

go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target ./gen

go generate ./ent

docker build -t sachinmv31/postgres-pgamber:latest .

docker-compose up -d

psql -h localhost -p 5432 -U postgres -d pgamber

CREATE EXTENSION plpython3u;

docker-compose down

go run main.go

docker volume rm setup_pgdata

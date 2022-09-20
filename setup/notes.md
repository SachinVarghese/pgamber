### Steps

go run -mod=mod entgo.io/ent/cmd/ent init Individual IncomeBracket

go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target ./gen

go generate ./ent

docker-compose up -d

psql -h localhost -p 5432 -U postgres -d pgamber

docker-compose down

go run main.go

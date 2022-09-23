### Setup Postgres Data

go run -mod=mod entgo.io/ent/cmd/ent init Individual IncomeBracket

go generate ./ent

go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target ./gen

go run main.go

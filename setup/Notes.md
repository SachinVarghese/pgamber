### Steps

go run -mod=mod entgo.io/ent/cmd/ent init Individual

go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target ./gen

go generate ./ent

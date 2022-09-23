package main

import (
	"context"
	"fmt"
	"log"

	ent_gen "github.com/SachinVarghese/pgamber/setup/ent/gen"
	"github.com/SachinVarghese/pgamber/setup/ingester"
	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "postgres"
	host     = "localhost"
	port     = "5432"
	dbName   = "pgamber"
	sslMode  = "disable"
)

func main() {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)

	client, err := ent_gen.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := ingester.IngestAdultIncomeData(ctx, client); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}

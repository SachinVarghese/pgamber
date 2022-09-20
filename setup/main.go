package main

import (
	"context"
	"fmt"
	"log"

	ent_gen "github.com/SachinVarghese/pgamber/setup/ent/gen"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/individual"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	age int
}

func main() {

	client, err := ent_gen.Open("sqlite3", "file:pent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var person Person
	person.age = 30

	if _, err = CreateIndividual(ctx, client, &person); err != nil {
		log.Fatal(err)
	}

	if _, err = QueryIndividual(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func CreateIndividual(ctx context.Context, client *ent_gen.Client, person *Person) (*ent_gen.Individual, error) {

	u, err := client.Individual.
		Create().
		SetAge(person.age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)
	return u, nil
}

func QueryIndividual(ctx context.Context, client *ent_gen.Client) (*ent_gen.Individual, error) {
	u, err := client.Individual.
		Query().
		Where(individual.AgeGT(20)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)
	return u, nil
}

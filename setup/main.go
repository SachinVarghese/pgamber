package main

import (
	"context"
	"fmt"
	"log"

	ent_gen "github.com/SachinVarghese/pgamber/setup/ent/gen"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/incomebracket"
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
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	peopleData := fetchPeopleData(CSVDataFilepath, CSVTruthFilepath)

	individuals, err := createIndividualsData(ctx, client, peopleData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d individuals added ", len(individuals))

	_, err = createIncomeBracketData(ctx, client, peopleData, individuals)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d individual brackets added ", len(individuals))

	count, err := queryIndividualsCount(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d individuals count total", count)
}

func createIndividualsData(ctx context.Context, client *ent_gen.Client, people []Person) (persons []*ent_gen.Individual, err error) {
	bulk := make([]*ent_gen.IndividualCreate, len(people))
	for i, person := range people {

		bulk[i] = client.Individual.
			Create().
			SetAge(person.age).
			SetWorkclass(int(person.workclass)).
			SetEducation(int(person.education)).
			SetMaritalStatus(int(person.maritalStatus)).
			SetOccupation(int(person.occupation)).
			SetRelationship(int(person.relationship)).
			SetRace(int(person.race)).
			SetSex(int(person.sex)).
			SetCapitalGain(person.capitalGain).
			SetCapitalLoss(person.capitalLoss).
			SetHoursPerWeek(person.hoursPerWeek).
			SetCountry(int(person.country))
	}

	persons, err = client.Individual.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed bulk appending individuals: %w", err)
	}

	return persons, nil
}

func createIncomeBracketData(ctx context.Context, client *ent_gen.Client, people []Person, persons []*ent_gen.Individual) (brackets []*ent_gen.IncomeBracket, err error) {
	bulk := make([]*ent_gen.IncomeBracketCreate, len(people))
	for i, person := range people {
		if person.class {
			bulk[i] = client.IncomeBracket.Create().SetClass(incomebracket.ClassGt50K).SetPersonID(persons[i].ID)
		} else {
			bulk[i] = client.IncomeBracket.Create().SetClass(incomebracket.ClassLte50K).SetPersonID(persons[i].ID)
		}
	}

	brackets, err = client.IncomeBracket.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed bulk appending individual income brackets: %w", err)
	}

	return brackets, nil
}

func queryIndividualsCount(ctx context.Context, client *ent_gen.Client) (int, error) {
	count, err := client.Individual.
		Query().Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed querying individuals: %w", err)
	}

	return count, nil
}

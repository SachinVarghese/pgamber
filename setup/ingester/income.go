package ingester

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ent_gen "github.com/SachinVarghese/pgamber/setup/ent/gen"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/incomebracket"
)

const (
	CSVDataFilepath  = "./data/array.csv"
	CSVTruthFilepath = "./data/truth.csv"
	CSVREcordSize    = 100
)

type Person struct {
	Age           float64
	Workclass     int64
	Education     int64
	MaritalStatus int64
	Occupation    int64
	Relationship  int64
	Race          int64
	Sex           int64
	CapitalGain   float64
	CapitalLoss   float64
	HoursPerWeek  float64
	Country       int64
	Class         bool
}

func IngestAdultIncomeData(ctx context.Context, client *ent_gen.Client) error {
	peopleData := fetchPeopleData()

	individuals, err := createIndividualsData(ctx, client, peopleData)
	if err != nil {
		return err
	}
	log.Printf("%d individuals added ", len(individuals))

	_, err = createIncomeBracketData(ctx, client, peopleData, individuals)
	if err != nil {
		return err
	}
	log.Printf("%d individual brackets added ", len(individuals))

	count, err := queryIndividualsCount(ctx, client)
	if err != nil {
		return err
	}
	log.Printf("%d individuals count total", count)
	return nil
}

func createIndividualsData(ctx context.Context, client *ent_gen.Client, people []Person) (persons []*ent_gen.Individual, err error) {
	bulk := make([]*ent_gen.IndividualCreate, len(people))
	for i, person := range people {

		bulk[i] = client.Individual.
			Create().
			SetAge(person.Age).
			SetWorkclass(int(person.Workclass)).
			SetEducation(int(person.Education)).
			SetMaritalStatus(int(person.MaritalStatus)).
			SetOccupation(int(person.Occupation)).
			SetRelationship(int(person.Relationship)).
			SetRace(int(person.Race)).
			SetSex(int(person.Sex)).
			SetCapitalGain(person.CapitalGain).
			SetCapitalLoss(person.CapitalLoss).
			SetHoursPerWeek(person.HoursPerWeek).
			SetCountry(int(person.Country))
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
		if person.Class {
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

func fetchPeopleData() (people []Person) {
	records := readCsvFile(CSVDataFilepath)
	classes := readCsvFile(CSVTruthFilepath)

	records = records[0:CSVREcordSize]
	classes = classes[0:CSVREcordSize]

	people = []Person{}
	for i, r := range records {
		person := Person{}

		v := r[0]
		v = strings.TrimPrefix(v, "[")
		v = strings.TrimSuffix(v, "]")
		values := strings.Split(v, " ")

		c := classes[i][0]
		c = strings.TrimPrefix(c, "[")
		c = strings.TrimPrefix(c, "[")
		classNum, err := strconv.ParseInt(c, 10, 64)

		age, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			break
		}
		person.Age = age

		workclass, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			break
		}
		person.Workclass = workclass

		education, err := strconv.ParseInt(values[2], 10, 64)
		if err != nil {
			break
		}
		person.Education = education

		maritalStatus, err := strconv.ParseInt(values[3], 10, 64)
		if err != nil {
			break
		}
		person.MaritalStatus = maritalStatus

		occupation, err := strconv.ParseInt(values[4], 10, 64)
		if err != nil {
			break
		}
		person.Occupation = occupation

		relationship, err := strconv.ParseInt(values[5], 10, 64)
		if err != nil {
			break
		}
		person.Relationship = relationship

		race, err := strconv.ParseInt(values[6], 10, 64)
		if err != nil {
			break
		}
		person.Race = race

		sex, err := strconv.ParseInt(values[7], 10, 64)
		if err != nil {
			break
		}
		person.Sex = sex

		capitalGain, err := strconv.ParseFloat(values[8], 64)
		if err != nil {
			break
		}
		person.CapitalGain = capitalGain

		capitalLoss, err := strconv.ParseFloat(values[9], 64)
		if err != nil {
			break
		}
		person.CapitalLoss = capitalLoss

		hoursPerWeek, err := strconv.ParseFloat(values[10], 64)
		if err != nil {
			break
		}
		person.HoursPerWeek = hoursPerWeek

		country, err := strconv.ParseInt(values[11], 10, 64)
		if err != nil {
			break
		}
		person.Country = country

		class := classNum > 0
		if err != nil {
			break
		}
		person.Class = class

		people = append(people, person)
	}
	return people
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

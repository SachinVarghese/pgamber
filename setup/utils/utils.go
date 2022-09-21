package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	CSVDataFilepath  = "./data/array.csv"
	CSVTruthFilepath = "./data/truth.csv"
	CSVREcordSize    = 1
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

func FetchPeopleData() (people []Person) {
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

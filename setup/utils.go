package main

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
	CSVREcordSize    = 1000
)

type Person struct {
	age           float64
	workclass     int64
	education     int64
	maritalStatus int64
	occupation    int64
	relationship  int64
	race          int64
	sex           int64
	capitalGain   float64
	capitalLoss   float64
	hoursPerWeek  float64
	country       int64
	class         bool
}

func fetchPeopleData(dataPath string, truthPath string) (people []Person) {
	records := readCsvFile(dataPath)
	classes := readCsvFile(truthPath)

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
		person.age = age

		workclass, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			break
		}
		person.workclass = workclass

		education, err := strconv.ParseInt(values[2], 10, 64)
		if err != nil {
			break
		}
		person.education = education

		maritalStatus, err := strconv.ParseInt(values[3], 10, 64)
		if err != nil {
			break
		}
		person.maritalStatus = maritalStatus

		occupation, err := strconv.ParseInt(values[4], 10, 64)
		if err != nil {
			break
		}
		person.occupation = occupation

		relationship, err := strconv.ParseInt(values[5], 10, 64)
		if err != nil {
			break
		}
		person.relationship = relationship

		race, err := strconv.ParseInt(values[6], 10, 64)
		if err != nil {
			break
		}
		person.race = race

		sex, err := strconv.ParseInt(values[7], 10, 64)
		if err != nil {
			break
		}
		person.sex = sex

		capitalGain, err := strconv.ParseFloat(values[8], 64)
		if err != nil {
			break
		}
		person.capitalGain = capitalGain

		capitalLoss, err := strconv.ParseFloat(values[9], 64)
		if err != nil {
			break
		}
		person.capitalLoss = capitalLoss

		hoursPerWeek, err := strconv.ParseFloat(values[10], 64)
		if err != nil {
			break
		}
		person.hoursPerWeek = hoursPerWeek

		country, err := strconv.ParseInt(values[11], 10, 64)
		if err != nil {
			break
		}
		person.country = country

		class := classNum > 0
		if err != nil {
			break
		}
		person.class = class

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

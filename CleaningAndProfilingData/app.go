package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	minScore = 640.0
	maxScore = 830.0
)

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func main() {
	f, err := os.Open("../data/loan_data.csv")
	must(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	rawCSVData, err := reader.ReadAll()
	must(err)

	f, err = os.Create("../data/clean_loan_data.csv")
	must(err)
	w := csv.NewWriter(f)

	for idx, record := range rawCSVData {
		if idx == 0 {
			must(w.Write(record))
			continue
		}
		encodedRecord := make([]string, 2)

		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		must(err)
		encodedRecord[0] = strconv.FormatFloat((score-minScore)/(maxScore-minScore), 'f', 4, 64)
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		must(err)

		if rate <= 12 {
			encodedRecord[1] = "1.0"
			must(w.Write(encodedRecord))
			continue
		}
		encodedRecord[1] = "0.0"
		must(w.Write(encodedRecord))
	}
	w.Flush()
	must(w.Error())
}

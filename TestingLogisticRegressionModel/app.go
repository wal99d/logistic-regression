package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func predict(score float64) float64 {
	p := 1 / (1 + math.Exp(-13.65*score+4.89))
	if p >= 0.5 {
		return 1.0
	}
	return 0.0
}

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func main() {
	f, err := os.Open("../data/testing.csv")
	must(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2
	records, err := reader.ReadAll()
	must(err)

	var observed []float64
	var predicted []float64

	for idx, record := range records {
		if idx == 0 {
			continue
		}
		score, err := strconv.ParseFloat(record[0], 64)
		must(err)
		observedVal, err := strconv.ParseFloat(record[1], 64)
		must(err)
		predictVal := predict(score)

		observed = append(observed, observedVal)
		predicted = append(predicted, predictVal)
	}
	var truePosNeg int

	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	accuracy := float64(truePosNeg) / float64(len(observed))

	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

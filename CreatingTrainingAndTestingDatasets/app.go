package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func main() {
	f, err := os.Open("../data/clean_loan_data.csv")
	must(err)
	defer f.Close()

	df := dataframe.ReadCSV(f)

	trainingNum := (4 * df.Nrow()) / 5
	testingNum := df.Nrow() / 5
	if trainingNum+testingNum < df.Nrow() {
		trainingNum++
	}
	trainingIdx := make([]int, trainingNum)
	testingIdx := make([]int, testingNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	for i := 0; i < testingNum; i++ {
		testingIdx[i] = testingNum + i
	}

	trainingDf := df.Subset(trainingIdx)
	testingDf := df.Subset(testingIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDf,
		1: testingDf,
	}

	for idx, csvFile := range []string{"training.csv", "testing.csv"} {
		f, err := os.Create("../data/" + csvFile)
		must(err)
		w := bufio.NewWriter(f)
		must(setMap[idx].WriteCSV(w))
	}
}

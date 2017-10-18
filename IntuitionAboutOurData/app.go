package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"

	"gonum.org/v1/plot/plotter"

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
	fmt.Printf("%v\n", df.Describe())
	for _, colName := range df.Names() {
		histVal := make(plotter.Values, df.Nrow())
		for idx, floatVal := range df.Col(colName).Float() {
			histVal[idx] = floatVal
		}

		p, err := plot.New()
		must(err)
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		h, err := plotter.NewHist(histVal, 16)
		must(err)
		h.Normalize(1)
		p.Add(h)

		must(p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"))
	}
}

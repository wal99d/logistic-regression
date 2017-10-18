package main

import (
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
)

func logisticFunc(x float64) float64 { return 1 / (1 + math.Exp(-x)) }

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
func main() {
	x := make(plotter.XYs, 100)
	for i := 0; i < len(x); i++ {
		x[i].X = float64(i)
		x[i].Y = logisticFunc(float64(-i))
	}

	p, err := plot.New()
	must(err)
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	logisticPlotter, err := plotter.NewScatter(x)
	must(err)
	logisticPlotter.Color = color.RGBA{B: 255, A: 255}
	p.Add(logisticPlotter)
	p.X.Max = 10
	p.X.Min = -10
	p.Y.Max = 1.1
	p.Y.Min = -0.1

	must(p.Save(4*vg.Inch, 4*vg.Inch, "logistic.png"))
}

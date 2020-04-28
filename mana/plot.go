package main

import (
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func makePlot(data map[int]int, title, filename string) {

	// sort manaRank
	keys := make([]int, len(data))
	i := 0
	for key := range data {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	points := []float64{}
	xy := make(plotter.XYs, len(keys))
	for i, key := range keys {
		points = append(points, float64(key))
		xy[i].X = float64(key)
		xy[i].Y = float64(data[key])
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title

	// Create a histogram of our values
	h, err := plotter.NewHistogram(xy, xy.Len())
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}

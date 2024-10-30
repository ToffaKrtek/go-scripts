package main

import (
	// "bytes"
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func makeDashBord(data []float64, labels []string) error {
	p := plot.New()
	p.Title.Text = "TEST"
	p.X.Label.Text = "TEST"
	p.X.Label.Text = "TEST"
	bars, err := plotter.NewBarChart(plotter.Values(data), vg.Points(20))
	if err != nil {
		return err
	}
	bars.Color = color.RGBA{R: 255, A: 255} // Красный цвет
	p.Add(bars)

	file, err := os.Create("plot.png")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := p.Save(4*vg.Inch, 4*vg.Inch, file.Name()); err != nil {
		return err
	}

	fmt.Println("График успешно сохранен в файл plot.png!")
	// Сохранение графика в буфер
	/*var buf bytes.Buffer
	w, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		return err
	}
	if _, err := w.WriteTo(&buf); err != nil {
		return err
	}*/
	return nil
}

func main() {
	makeDashBord([]float64{6, 4, 1, 7}, []string{"Ras", "Dva", "TREE", "chertary"})
}

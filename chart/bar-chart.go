package chart

import (
	"math"

	"github.com/spiegel-im-spiegel/errs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func MakeHistChart(list *HistList, title, outPath string) error {
	labelX := []string{}
	dataY1 := plotter.Values{}
	dataY2 := plotter.Values{}
	dataY3 := plotter.XYs{}
	maxCases := 0.0
	for i := 0; i < list.Size(); i++ {
		d := list.Data(i)
		labelX = append(labelX, d.Period.StringEnd())
		dataY1 = append(dataY1, math.Floor(d.Cases))
		maxCases = max(maxCases, d.Cases)
		dataY2 = append(dataY2, math.Floor(d.Deaths))
		maxCases = max(maxCases, d.Deaths)
		dataY3 = append(dataY3, plotter.XY{X: (float64)(i), Y: math.Floor(d.Hospitalized)})
		maxCases = max(maxCases, d.Hospitalized)
	}
	maxCases = maxCases * 5 / 3
	maxCases = (float64)((((int)(maxCases) / 50) + 1) * 50)

	//default font
	plot.DefaultFont = "Helvetica"
	plotter.DefaultFont = "Helvetica"

	//new plot
	p, err := plot.New()
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}

	//new bar chart
	bar1, err := plotter.NewBarChart(dataY1, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar2, err := plotter.NewBarChart(dataY2, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}

	bar1.LineStyle.Width = vg.Length(0)
	bar1.Color = plotutil.Color(2)
	bar1.Offset = -2
	bar1.Horizontal = false
	p.Add(bar1)

	bar2.LineStyle.Width = vg.Length(0)
	bar2.Color = plotutil.Color(6)
	bar2.Offset = 2
	bar2.Horizontal = false
	p.Add(bar2)

	//new line chart
	line1, err := plotter.NewLine(dataY3)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}

	line1.Color = plotutil.Color(4)
	p.Add(line1)

	//labels of X
	p.NominalX(labelX...)
	p.X.Label.Text = "Date of report"
	//p.X.Padding = 0
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter

	//labels of Y
	p.Y.Label.Text = "Confirmed cases"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = math.Ceil(maxCases)

	//legend
	p.Legend.Add("New Cases", bar1)
	p.Legend.Add("New Deaths", bar2)
	p.Legend.Add("Hospitalized", line1)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = 0

	//title
	p.Title.Text = "Cases in " + title

	//output image
	if err := p.Save(23.0*(vg.Length)(list.Size()+2), 8*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	return nil
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

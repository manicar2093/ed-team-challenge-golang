package services

import (
	"bytes"
	"fmt"
	"io"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/manicar2093/ed-team-go/internal/stringsfmt"
	"github.com/manicar2093/ed-team-go/models"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

type ChartServiceImpl struct{}

func (c ChartServiceImpl) CreateCryptoChart(nomicsResponse []models.NomicsResponse) (io.WriterTo, error) {
	var pricesData []chart.Series

	var currencies []string

	for _, v := range nomicsResponse {
		randColor := randomcolor.GetRandomColorInRgb()
		color := drawing.Color{R: uint8(randColor.Red), G: uint8(randColor.Green), B: uint8(randColor.Blue), A: 255}
		temp := chart.TimeSeries{
			Name:    v.Currency,
			XValues: v.Timestamps,
			YValues: v.Prices,
			Style: chart.Style{
				Show:        true,
				StrokeColor: color,
				DotColor:    color,
			},
		}
		currencies = append(currencies, v.Currency)
		pricesData = append(pricesData, temp)
	}

	var currenciesToUse string
	if len(currencies) > 1 {
		currenciesToUse = c.createChartName(currencies)
	} else {
		currenciesToUse = currencies[0]
	}
	graphicName := fmt.Sprintf("%s Price(s) Graphic", currenciesToUse)

	graphic := chart.Chart{
		Title: graphicName,
		TitleStyle: chart.Style{
			Show:      true,
			FontSize:  15.0,
			FontColor: chart.ColorBlack,
			TextWrap:  chart.TextWrapWord,
		},
		XAxis: chart.XAxis{
			Style:        chart.StyleShow(),
			TickPosition: chart.TickPositionBetweenTicks,
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Range: c.createContinuousRangeY(nomicsResponse),
		},
		Series: pricesData,
	}
	graphic.Elements = []chart.Renderable{
		chart.Legend(&graphic),
	}

	graphFile := bytes.NewBuffer(nil)

	err := graphic.Render(chart.PNG, graphFile)
	if err != nil {
		return nil, err
	}
	return graphFile, nil
}

func (c ChartServiceImpl) createChartName(currenciesNames []string) string {
	if len(currenciesNames) == 2 {
		return fmt.Sprintf("%s and %s", currenciesNames[0], currenciesNames[1])
	}
	lastIndex := len(currenciesNames) - 1
	separated := stringsfmt.SeparateByCommas(currenciesNames[:lastIndex]...)
	return fmt.Sprintf("%s and %s", separated, currenciesNames[lastIndex])
}

func (c ChartServiceImpl) createContinuousRangeY(d []models.NomicsResponse) *chart.ContinuousRange {
	var min []float64
	var max []float64

	for _, v := range d {
		max = append(max, c.maxFloatSlice(v.Prices))
		min = append(min, c.minFloatSlice(v.Prices))
	}
	return &chart.ContinuousRange{
		Min: min[0],
		Max: max[0],
	}
}

func (c ChartServiceImpl) minFloatSlice(f []float64) (min float64) {
	min = f[0]
	for _, v := range f {
		if v < min {
			min = v
		}
	}
	return
}

func (c ChartServiceImpl) maxFloatSlice(f []float64) (max float64) {
	max = f[0]
	for _, v := range f {
		if v > max {
			max = v
		}
	}
	return
}

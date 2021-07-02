package services

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/manicar2093/ed-team-go/config"
	"github.com/manicar2093/ed-team-go/models"
	"github.com/manicar2093/ed-team-go/utils"
	"github.com/wcharczuk/go-chart"
)

type ChartService struct{}

func (c ChartService) CreateCryptoChart(d []models.NomicsResponse) (graphFile *os.File, e error) {
	var pricesData []chart.Series

	var currencies []string

	for i, v := range d {
		temp := chart.TimeSeries{
			Name:    v.Currency,
			XValues: v.Timestamps,
			YValues: v.Prices,
			Style: chart.Style{
				Show:        true,
				StrokeColor: chart.GetAlternateColor(i),
				DotColor:    chart.GetAlternateColor(i),
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
			Range: c.createContinuousRangeY(d),
		},
		Series: pricesData,
	}

	graphFile, e = os.Create(fmt.Sprintf("%s/%s_%s.png", config.FilesPath, time.Now().Format(time.RFC3339), graphicName))
	if e != nil {
		return
	}

	e = graphic.Render(chart.PNG, graphFile)
	if e != nil {
		return
	}
	return
}

func (c ChartService) createChartName(currenciesNames []string) string {
	if len(currenciesNames) == 2 {
		return fmt.Sprintf("%s and %s", currenciesNames[0], currenciesNames[1])
	}
	lastIndex := len(currenciesNames) - 2
	separated := utils.SeparateByCommas(currenciesNames[:lastIndex]...)
	return fmt.Sprintf("%s and %s", separated, currenciesNames[lastIndex+1])
}

func (c ChartService) createContinuousRangeY(d []models.NomicsResponse) *chart.ContinuousRange {
	var min []float64
	var max []float64

	for _, v := range d {
		max = append(max, c.maxFloatSlice(v.Prices))
		min = append(min, c.minFloatSlice(v.Prices))
	}
	totalMin := c.maxFloatSlice(min)
	totalMax := c.maxFloatSlice(max)
	return &chart.ContinuousRange{
		Min: totalMin,
		Max: totalMax,
	}
}

func (c ChartService) minFloatSlice(f []float64) float64 {
	sorted := sort.Float64Slice(sort.Float64Slice(f))
	return sorted[0]
}

func (c ChartService) maxFloatSlice(f []float64) float64 {
	sorted := sort.Float64Slice(sort.Float64Slice(f))
	return sorted[len(f)-1]
}

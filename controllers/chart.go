package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manicar2093/ed-team-go/internal/response"
	"github.com/manicar2093/ed-team-go/models"
	"github.com/manicar2093/ed-team-go/services"
)

type ChartController struct {
	chartService  services.ChartCreator
	nomicsService services.NomicsService
}

func NewChartController(chartService services.ChartCreator, nomicsService services.NomicsService) *ChartController {
	return &ChartController{chartService: chartService, nomicsService: nomicsService}
}

func (c ChartController) CreateChartHandler(w http.ResponseWriter, r *http.Request) {
	var charInfo models.CreateChartRequest
	e := json.NewDecoder(r.Body).Decode(&charInfo)
	if e != nil {
		log.Println(e)
		response.JSON(w, http.StatusBadRequest, map[string]string{"message": "request does not satisfy requirements. check documentation"})
		return
	}

	chartData, e := c.nomicsService.GetCryptoInfo(charInfo)
	if e != nil {
		log.Println(e)
		response.JSON(w, http.StatusInternalServerError, map[string]string{"message": "there was an error getting currencies data. please try later"})
		return
	}

	chartFile, e := c.chartService.CreateCryptoChart(chartData)
	if e != nil {
		log.Println(e)
		response.JSON(w, http.StatusInternalServerError, map[string]string{"message": "there was an error creating graph"})
		return
	}

	response.MakeDownloadableFile(w, chartFile, "currency-chart", response.ContentTypePNG)

}

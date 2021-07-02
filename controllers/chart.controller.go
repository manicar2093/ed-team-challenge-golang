package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manicar2093/ed-team-go/models"
	"github.com/manicar2093/ed-team-go/services"
	"github.com/manicar2093/ed-team-go/utils"
)

type ChartController struct {
	utils.ControllerUtils
	chartService  *services.ChartService
	nomicsService *services.NomicsService
}

func NewChartController(chartService *services.ChartService, nomicsService *services.NomicsService) *ChartController {
	return &ChartController{chartService: chartService, nomicsService: nomicsService}
}

func (c ChartController) CreateChartHandler(w http.ResponseWriter, r *http.Request) {
	var charInfo models.CreateChartRequest
	e := json.NewDecoder(r.Body).Decode(&charInfo)
	if e != nil {
		log.Println(e)
		c.JSON(w, http.StatusBadRequest, map[string]string{"message": "request does not satisfy requirements. check documentation"})
		return
	}

	chartData, e := c.nomicsService.GetCryptoInfo(charInfo)
	if e != nil {
		log.Println(e)
		c.JSON(w, http.StatusInternalServerError, map[string]string{"message": "there was an error getting currencies data. please try later"})
		return
	}

	chartFile, e := c.chartService.CreateCryptoChart(chartData)
	if e != nil {
		log.Println(e)
		c.JSON(w, http.StatusInternalServerError, map[string]string{"message": "there was an error creating graph"})
		return
	}

	c.MakeDownloadableFile(w, chartFile, utils.ContentTypePNG)

}

package services

import (
	"os"

	"github.com/manicar2093/ed-team-go/models"
)

type ChartService interface {
	CreateCryptoChart(nomicsResponse []models.NomicsResponse) (*os.File, error)
}

type NomicsService interface {
	GetCryptoInfo(ccr models.CreateChartRequest) ([]models.NomicsResponse, error)
}

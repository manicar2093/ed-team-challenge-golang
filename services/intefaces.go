package services

import (
	"io"

	"github.com/manicar2093/ed-team-go/models"
)

type ChartCreator interface {
	CreateCryptoChart(nomicsResponse []models.NomicsResponse) (io.WriterTo, error)
}

type NomicsService interface {
	GetCryptoInfo(ccr models.CreateChartRequest) ([]models.NomicsResponse, error)
}

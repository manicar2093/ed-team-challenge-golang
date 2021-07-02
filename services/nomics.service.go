package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/manicar2093/ed-team-go/config"
	"github.com/manicar2093/ed-team-go/models"
	"github.com/manicar2093/ed-team-go/utils"
)

type NomicsService struct {
	client *http.Client
}

func NewNomicsServiceWClient() *NomicsService {
	return &NomicsService{
		client: &http.Client{},
	}
}

func (n NomicsService) GetCryptoInfo(ccr models.CreateChartRequest) (data []models.NomicsResponse, e error) {
	req, e := n.createCurrenciesSparklineRequest(ccr)
	if e != nil {
		return
	}

	res, e := n.client.Do(req)
	if e != nil {
		return
	}

	e = json.NewDecoder(res.Body).Decode(&data)
	if e != nil {
		return
	}

	return
}

func (n NomicsService) createCurrenciesSparklineRequest(ccr models.CreateChartRequest) (r *http.Request, e error) {
	r, e = http.NewRequest(http.MethodGet, config.NominicsAPI, nil)
	if e != nil {
		return
	}
	cryptoIDs := utils.SeparateByCommas(ccr.Cryptos...)
	formatedDate := ccr.FromDate.Format(time.RFC3339)
	p := r.URL.Query()

	p.Set("ids", cryptoIDs)
	p.Set("key", config.NomicsKey)
	p.Set("start", formatedDate)
	r.URL.RawQuery = p.Encode()
	return
}

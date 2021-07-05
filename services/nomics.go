package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/manicar2093/ed-team-go/config"
	"github.com/manicar2093/ed-team-go/internal/stringsfmt"
	"github.com/manicar2093/ed-team-go/models"
)

type NomicsServiceImpl struct {
	endpoint string
	client   *http.Client
}

func NewNomicsServiceWClient(endpoint string) NomicsService {
	return &NomicsServiceImpl{
		endpoint: endpoint,
		client:   &http.Client{},
	}
}

func (n NomicsServiceImpl) GetCryptoInfo(ccr models.CreateChartRequest) ([]models.NomicsResponse, error) {
	req, e := n.createCurrenciesSparklineRequest(ccr)
	if e != nil {
		return []models.NomicsResponse{}, e
	}

	res, e := n.client.Do(req)
	if e != nil {
		return []models.NomicsResponse{}, e
	}

	var data []models.NomicsResponse
	e = json.NewDecoder(res.Body).Decode(&data)
	defer res.Body.Close()
	if e != nil {
		return []models.NomicsResponse{}, e
	}

	return data, nil
}

func (n NomicsServiceImpl) createCurrenciesSparklineRequest(ccr models.CreateChartRequest) (*http.Request, error) {
	r, e := http.NewRequest(http.MethodGet, n.endpoint, nil)
	if e != nil {
		return nil, e
	}
	cryptoIDs := stringsfmt.SeparateByCommas(ccr.Cryptos...)
	formatedDate := ccr.FromDate.Format(time.RFC3339)
	p := r.URL.Query()

	p.Set("ids", cryptoIDs)
	p.Set("key", config.NomicsKey)
	p.Set("start", formatedDate)
	r.URL.RawQuery = p.Encode()
	return r, e
}

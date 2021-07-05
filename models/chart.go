package models

import "time"

type CreateChartRequest struct {
	Cryptos  []string
	FromDate time.Time `json:"from_date"`
}

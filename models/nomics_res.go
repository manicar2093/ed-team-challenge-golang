package models

import (
	"encoding/json"
	"strconv"
	"time"
)

type NomicsResponse struct {
	Currency   string
	Timestamps []time.Time
	Prices     []float64
}

func (f *NomicsResponse) UnmarshalJSON(data []byte) (e error) {
	var values map[string]interface{}

	if e = json.Unmarshal(data, &values); e != nil {
		return
	}

	f.Currency = values["currency"].(string)
	var prices []float64

	for _, v := range values["prices"].([]interface{}) {
		priceFloat, e := strconv.ParseFloat(v.(string), 64)
		if e != nil {
			break
		}
		prices = append(prices, priceFloat)
	}

	f.Prices = prices

	var timestamps []time.Time

	for _, v := range values["timestamps"].([]interface{}) {
		tempTime, e := time.Parse(time.RFC3339, v.(string))
		if e != nil {
			break
		}
		timestamps = append(timestamps, tempTime)
	}

	f.Timestamps = timestamps

	return e
}

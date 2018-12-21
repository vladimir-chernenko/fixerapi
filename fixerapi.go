package fixerapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type CurrencyRates struct {
	Success   bool               `json:"success"`
	Timestamp int32              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float32 `json:"rates"`
}

type FixerClient struct {
	apiKey string
	client http.Client
}

func NewFixerClient(apiKey string) *FixerClient {
	return &FixerClient{
		apiKey: apiKey,
		client: http.Client{Timeout: time.Second * 2},
	}
}

func (fc *FixerClient) GetCurrencyRates() (CurrencyRates, error) {
	var cr CurrencyRates
	baseURL := "http://data.fixer.io/api/latest?access_key=" + fc.apiKey

	r, err := fc.client.Get(baseURL)
	if err != nil {
		return cr, err
	}

	err = json.NewDecoder(r.Body).Decode(&cr)
	return cr, err
}

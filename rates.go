package gopenexchangerates

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var Empty = errors.New("Cache empty, ensure you've successfully called Populate at least once.")

type ExchangeRates struct {
	Base      string
	rates     map[string]float64
	Timestamp time.Time
	url       string
}

// Create a new exchange rates cache
func New(appid string) *ExchangeRates {
	return &ExchangeRates{url: "http://openexchangerates.org/latest.json?app_id=" + appid}
}

type deserialised struct {
	Base      string
	Rates     map[string]float64
	Timestamp int64
}

// Populate the rates from Open Exchange Rates
func (r *ExchangeRates) Populate() error {
	res, err := http.Get(r.url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var d deserialised
	if err := json.Unmarshal(body, &d); err != nil {
		return err
	}

	r.Base = d.Base
	r.Timestamp = time.Unix(d.Timestamp, 0)
	r.rates = d.Rates

	return nil
}

// Get all the exchange rates
func (r *ExchangeRates) All() (map[string]float64, error) {
	if r.rates == nil {
		return nil, Empty
	}

	rates := make(map[string]float64, len(r.rates))
	for k, v := range r.rates {
		rates[k] = v
	}

	return rates, nil
}

// Get the exchange rate for a currency
func (r *ExchangeRates) Get(currency string) (float64, error) {
	if r.rates == nil {
		return 0, Empty
	}

	return r.rates[currency], nil
}

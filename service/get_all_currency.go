package service

import (
	"crypto-server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllCurrency() model.AllCurrencyResponse {
	fmt.Println("Calling API...")

	currencies := GetAllCurrencyURL()
	symbols := GetAllSymbolsURL()
	tickers := GetAllTickersURL()

	var data model.AllCurrencyResponse
	for a, b := range symbols {
		var res model.CurrencyResponse

		res.ID = b.BaseCurrency

		res.FullName = currencies[res.ID].FullName
		res.Ask = tickers[a].Ask
		res.Bid = tickers[a].Bid
		res.High = tickers[a].High
		res.Last = tickers[a].Last
		res.Low = tickers[a].Low
		res.FeeCurrency = b.FeeCurrency

		data.Currencies = append(data.Currencies, res)
	}

	return data
}

func GetAllCurrencyURL() map[string]model.Currency {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/currency", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	data := make(map[string]model.Currency)

	json.Unmarshal(bodyBytes, &data)

	return data
}

func GetAllSymbolsURL() map[string]model.Symbol {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/symbol", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	data := make(map[string]model.Symbol)

	json.Unmarshal(bodyBytes, &data)

	return data
}

func GetAllTickersURL() map[string]model.Ticker {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/ticker", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	data := make(map[string]model.Ticker)

	json.Unmarshal(bodyBytes, &data)

	return data
}

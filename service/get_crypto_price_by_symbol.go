package service

import (
	"crypto-server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllCryptoPriceBySymbol(symbolName string) model.CurrencyResponse {
	fmt.Println("Calling API...")

	symbol := GetSymbolsURL(symbolName)
	ticker := GetTickersURL(symbolName)
	currency := GetCurrencyURL(symbol.BaseCurrency)

	var res model.CurrencyResponse

	res.ID = symbol.BaseCurrency

	res.FullName = currency.FullName
	res.Ask = ticker.Ask
	res.Bid = ticker.Bid
	res.High = ticker.High
	res.Last = ticker.Last
	res.Low = ticker.Low
	res.FeeCurrency = symbol.FeeCurrency

	return res
}

func GetCurrencyURL(name string) model.Currency {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/currency/"+name, nil)
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

	var data model.Currency

	json.Unmarshal(bodyBytes, &data)

	return data
}

func GetSymbolsURL(name string) model.Symbol {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/symbol/"+name, nil)
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

	var data model.Symbol
	json.Unmarshal(bodyBytes, &data)

	return data
}

func GetTickersURL(name string) model.Ticker {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/3/public/ticker/"+name, nil)
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

	var data model.Ticker

	json.Unmarshal(bodyBytes, &data)

	return data
}

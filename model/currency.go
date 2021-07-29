package model

type Currency struct {
	FullName string `json:"full_name"`
}

type Symbol struct {
	FeeCurrency  string `json:"fee_currency"`
	BaseCurrency string `json:"base_currency"`
}

type Ticker struct {
	Ask  string `json:"ask"`
	Bid  string `json:"bid"`
	Last string `json:"last"`
	Open string `json:"open"`
	Low  string `json:"low"`
	High string `json:"high"`
}

type CurrencyResponse struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"fee_currency"`
}

type AllCurrencyResponse struct {
	Currencies []CurrencyResponse `json:"currencies"`
}

package main

//Configuration hold the main config.json content
type Configuration struct {
	Currency         []Currency `json:"btc_exchange"`
	CurrencyExchange string     `json:"exchange"`
	OriginCurrency   string     `json:"active_origin_currency"`
	DestinyCurrency  string     `json:"active_destiny_currency"`
}

//Currency holds the main struct for the currency and its exchange api
type Currency struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//ConversionOperation holds the data needed for an operation
type ConversionOperation struct {
	originCurrency       float64
	destinyCurrency      float64
	originCurrencyToBtc  float64
	destinyCurrencyToBtc float64
}

//ExchangeCurrencyAPI Currency Exchange struct for API response
type ExchangeCurrencyAPI struct {
	Base     string   `json:"base"`
	USARates USARates `json:"rates"`
}

//USARates Currency Exchange struct for API response
type USARates struct {
	MXNRate float64 `json:"MXN"`
}

//MxnToBTC Handles the BTC to MXN json response
type MxnToBTC struct {
	Success bool            `json:"success"`
	Payload MxnToBTCPayload `json:"payload"`
}

//MxnToBTCPayload Handles part of the BTC to MXN json response
type MxnToBTCPayload struct {
	Bid  string `json:"bid"`
	Ask  string `json:"ask"`
	High string `json:"high"`
	Low  string `json:"low"`
}

//USDToBTC Handles the BTC to USD json response
type USDToBTC struct {
	//Currency float64 `json:"bpi.USD.rate_float"`
	Bpi BpiUSDToBTC `json:"bpi"`
}

//BpiUSDToBTC Handles the BTC to USD json response
type BpiUSDToBTC struct {
	USD USDUSDToBTC `json:"usd"`
}

//USDUSDToBTC Handles the BTC to USD json response
type USDUSDToBTC struct {
	RateFloat float64 `json:"rate_float"`
}

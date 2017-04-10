package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//GetLeatestExchangeRates Pull information from all exchange URL's
func GetLeatestExchangeRates() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	//Get BTC to USD
	getLeatestBtcToUSDRate(wg)

	//Get BTC to MXN
	getLeatestBtcToMxnRate(wg)

	//Get MXN to USD
	getLeatestCurrencyRate(wg)

	wg.Wait()
}

func getLeatestCurrencyRate(wg *sync.WaitGroup) {
	if *mxntousd > 0 {
		conversionOperation.destinyCurrency = *mxntousd
		wg.Done()
		return
	}

	var exchangeCurrencyAPI = new(ExchangeCurrencyAPI)
	getJSON(config.CurrencyExchange, exchangeCurrencyAPI)

	conversionOperation.destinyCurrency = exchangeCurrencyAPI.USARates.MXNRate

	wg.Done()
	fmt.Println("Currency exchange Rate Updated")
}

func getLeatestBtcToMxnRate(wg *sync.WaitGroup) {
	var mxnToBTC = new(MxnToBTC)
	mxnURL := getURLFromCurrency(config.DestinyCurrency)

	getJSON(mxnURL, mxnToBTC)

	f, err := strconv.ParseFloat(mxnToBTC.Payload.Bid, 64)

	if err != nil {
		log.Panic(err)
	}

	conversionOperation.destinyCurrencyToBtc = f

	wg.Done()
	fmt.Println("BTC to MXN Rate Updated")
}

func getLeatestBtcToUSDRate(wg *sync.WaitGroup) {
	var usdToBTC = new(USDToBTC)
	usdURL := getURLFromCurrency(config.OriginCurrency)

	getJSON(usdURL, usdToBTC)

	conversionOperation.originCurrencyToBtc = usdToBTC.Bpi.USD.RateFloat

	wg.Done()
	fmt.Println("BTC to USD Rate Updated")
}

//getJSON Retrieves a json file from the given URL and parse it to the given struct
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//getURLFromCurrency Search in the array of currencys URL's for th given currency
func getURLFromCurrency(currency string) string {
	for _, v := range config.Currency {
		if v.Name == currency {
			return v.URL
		}
	}
	return ""
}

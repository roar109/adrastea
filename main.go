package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	config              Configuration
	conversionOperation = ConversionOperation{}
	usdToBuy            = flag.Float64("usd", 1000.0, "USD to buy")
	mxntousd            = flag.Float64("mxntousd", 0.0, "Overwrite MXN to USD exchange")
)

func main() {
	//Parse any config
	flag.Parse()
	config = *readConfigFile()
	fmt.Println("Origin currency=" + config.OriginCurrency)
	fmt.Println("Destiny currency=" + config.DestinyCurrency)
	fmt.Println(" ")

	//Pull exchange rates
	//TODO do this each 5min in a future version
	GetLeatestExchangeRates()
	fmt.Println(" ")

	//Operations parse
	conversionOperation.originCurrency = *usdToBuy

	originConv := conversionOperation.getMXNWithCurrentUSD()
	fmt.Println("MXN to pay:", originConv)

	btcFromOrigin := conversionOperation.getBTCFromOrigin()
	fmt.Println("BTC with", *usdToBuy, "usd:", btcFromOrigin)

	destinyTotal := conversionOperation.getDestinyFromBTC(btcFromOrigin)
	fmt.Println(" ")
	fmt.Println("Final amount to withdraw (MXN):", destinyTotal)
	fmt.Println("Total earnings (initial  - current, without commission):", destinyTotal-originConv)
}

func readConfigFile() *Configuration {
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsontype Configuration
	err := json.Unmarshal(file, &jsontype)
	if err != nil {
		log.Fatal(err)
	}
	return &jsontype
}

func (co ConversionOperation) getMXNWithCurrentUSD() float64 {
	fmt.Println("USD to buy", co.originCurrency, "with a rate of", co.destinyCurrency, "MXN per USD")
	return co.destinyCurrency * co.originCurrency
}

func (co ConversionOperation) getBTCFromOrigin() float64 {
	fmt.Println("Transforming ", *usdToBuy, "USD into BTC at rate of", co.originCurrencyToBtc, "USD per BTC")
	return *usdToBuy / co.originCurrencyToBtc
}

func (co ConversionOperation) getDestinyFromBTC(btc float64) float64 {
	fmt.Println("Transforming", btc, "BTC into MXN at a rate of", co.destinyCurrencyToBtc, "MXN per BTC")
	return btc * co.destinyCurrencyToBtc
}

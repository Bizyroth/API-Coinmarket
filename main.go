package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	//"os"
	"io/ioutil"
	"net/http"
	"time"

)

type aPI struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	Rank            string  `json:"rank"`
	PriceUsd        float64 `json:"price_usd"`
	PriceBtc        string  `json:"price_btc"`
	VolUsd          float64 `json:"24h_volume_usd"`
	MarketCapUsd    float64 `json:"market_cap_usd"`
	AvailableSupply float64 `json:"available_supply"`
	TotalSupply     float64 `json:"total_supply"`
	PercentChange1h float64 `json:"percent_change_1h"`
	Percent24h      string  `json:"percent_change_24h"`
	Percent7d       string  `json:"percent_change_7d"`
	LastUpdated     float64 `json:"last_updated"`
	PriceEur        string  `json:"price_eur"`
	VolEur          float64 `json:"24h_volume_eur"`
	MarketCapEur    float64 `json:"market_cap_eur"`
}

func main() {
	//i := 0
	var aa []aPI
	for {
		url := "https://api.coinmarketcap.com/v1/ticker/eos/?convert=EUR"
		cli := http.Client{}

		req, err := cli.Get(url)
		if err != nil {
			fmt.Println("newrequest " + err.Error())
		}
		defer req.Body.Close()
		
		body, err := ioutil.ReadAll(req.Body)

		json.Unmarshal(body, &aa)
		eosPriceBt, _ := strconv.ParseFloat(aa[0].PriceBtc, 64)
		fmt.Printf("price eos: %v\n", eosPriceBt)
		
		url = "https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=EUR"
		req, err = cli.Get(url)
		if err != nil {
			fmt.Println("newrequest " + err.Error())
		}
		defer req.Body.Close()
		
		body, err = ioutil.ReadAll(req.Body)

		json.Unmarshal(body, &aa)
		ethereumPriceBt, _ := strconv.ParseFloat(aa[0].PriceBtc, 64)		
		fmt.Printf("price ethereum: %v\n", ethereumPriceBt)
		

		/*if(aa[l].ID=="eos" || aa[l].ID=="ethereum"){
			//fmt.Println(string(body))
			priceEU, _ := strconv.ParseFloat(aa[l].PriceEur, 64)
			percent24, _ := strconv.ParseFloat(aa[l].Percent24h, 64)
			percent7, _ := strconv.ParseFloat(aa[l].Percent7d, 64)
			priceBt, _ := strconv.ParseFloat(aa[l].PriceBtc, 64)
			
			fmt.Printf("%20v     %5v      %15v       %6v      %10v      %10v", aa[l].ID, aa[l].Symbol, priceEU, percent24, percent7, priceBt)
			fmt.Println()
			
			
			
			}
		}
*/		fmt.Printf("Price EOS in ethereum: %v\n",eosPriceBt/ethereumPriceBt)
		time.Sleep(5 * time.Minute)
		}
}


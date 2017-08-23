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
	url := "https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=10"

	for {
		cli := http.Client{}

		req, err := cli.Get(url)
		if err != nil {
			fmt.Println("newrequest " + err.Error())
		}
		defer req.Body.Close()

		body, err := ioutil.ReadAll(req.Body)
		var aa []aPI
		json.Unmarshal(body, &aa)
		fmt.Printf("%20s     %5s      %15s       %6s      %10s      %10s", "ID", "symbol", "price(eur)", "%24", "%7", "price (btc)")
		fmt.Println()
		for l := range aa {

			//fmt.Println(string(body))
			priceEU, _ := strconv.ParseFloat(aa[l].PriceEur, 64)
			percent24, _ := strconv.ParseFloat(aa[l].Percent24h, 64)
			percent7, _ := strconv.ParseFloat(aa[l].Percent7d, 64)
			priceBt, _ := strconv.ParseFloat(aa[l].PriceBtc, 64)

			fmt.Printf("%20v     %5v      %15v       %6v      %10v      %10v", aa[l].ID, aa[l].Symbol, priceEU, percent24, percent7, priceBt)
			fmt.Println()
		}
		time.Sleep(5 * time.Minute)
	}

}

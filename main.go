package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type aPI struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Rank            string `json:"rank"`
	PriceUsd        string `json:"price_usd"`
	PriceBtc        string `json:"price_btc"`
	VolUsd          string `json:"24h_volume_usd"`
	MarketCapUsd    string `json:"market_cap_usd"`
	AvailableSupply string `json:"available_supply"`
	TotalSupply     string `json:"total_supply"`
	PercentChange1h string `json:"percent_change_1h"`
	Percent24h      string `json:"percent_change_24h"`
	Percent7d       string `json:"percent_change_7d"`
	LastUpdated     string `json:"last_updated"`
	PriceEur        string `json:"price_eur"`
	VolEur          string `json:"24h_volume_eur"`
	MarketCapEur    string `json:"market_cap_eur"`
}

func getAPICoinmarket(money string) (ap []aPI, err error) {
	var a []aPI
	url := "https://api.coinmarketcap.com/v1/ticker/" + money + "/?convert=EUR"
	cli := http.Client{}
	req, err := cli.Get(url)
	if err != nil {
		fmt.Println("newrequest " + err.Error())
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &a)
	return a, err
}

func getDouble(num string) (float64, error) {
	return strconv.ParseFloat(num, 64)
}

type eosAPI struct {
	ID          int64   `json:"id"`
	CreateOnDay float64 `json:"createOnDay"`
	DailyTotal  float64 `json:"dailyTotal"`
	Price       float64 `json:"Price"`
	Begins      string  `json:"begins"`
	Ends        string  `json:"ends"`
}

func getEosAPI() (ap []eosAPI, err error) {
	var a []eosAPI
	url := "https://eos.io/eos-sales-statistic.php"
	cli := http.Client{}
	req, err := cli.Get(url)
	if err != nil {
		fmt.Println("newrequest " + err.Error())
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &a)
	return a, err
}

func compareDate(t1 time.Time, t2 time.Time) int {
	if t1.Year() < t2.Year() {
		return 1
	}
	if t1.Year() > t2.Year() {
		return -1
	}
	if t1.Month() < t2.Month() {
		return 1
	}
	if t1.Month() > t2.Month() {
		return -1
	}
	if t1.Day() < t2.Day() {
		return 1
	}
	if t1.Day() > t2.Day() {
		return -1
	}
	if t1.Minute() < t2.Minute() {
		return 1
	}
	if t1.Minute() > t2.Minute() {
		return -1
	}
	if t1.Second() < t2.Second() {
		return 1
	}
	if t1.Second() > t2.Second() {
		return -1
	}
	return 0
}

func getPeriod(ap []eosAPI) int {
	i := 0
	t := time.Now().UTC()
	for i < 365 {
		tB, _ := time.Parse(time.RFC3339, ap[i].Begins)
		tE, _ := time.Parse(time.RFC3339, ap[i].Ends)
		if compareDate(tB, t) == 1 && compareDate(t, tE) == 1 {
			return i
		}
		i++
	}
	return -1
}

func getNumberToken(totalOwnEth float64, totalEosToken float64, totalAllEth float64) float64 {
	return totalOwnEth * (totalEosToken / totalAllEth)
}

func mustBuy(numToken float64, priceEos float64, priceEth float64) bool {
	priceEosInEth := 1 / numToken

	priceEosSell := priceEosInEth * priceEth

	fmt.Printf("Price Eos to buy: \t\t\t%v\n", priceEosSell)
	fmt.Printf("Price Eos to sell: \t\t\t%v\n", priceEos)
	if priceEos > priceEosSell {
		fmt.Printf("gain en pourcentage \t\t\t%v\n", 100*priceEos/priceEosSell)
		return true
	}
	return false
}

func main() {

	for {
		eosAP, err := getAPICoinmarket("eos")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		ethAP, err := getAPICoinmarket("ethereum")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		ethPriceUs, err := getDouble(ethAP[0].PriceUsd)
		eosPriceUs, err := getDouble(eosAP[0].PriceUsd)
		if err != nil {
			fmt.Println("Error conversion: " + err.Error())
		}

		fmt.Printf("Price EOS in USD: \t\t\t%v\n", eosPriceUs)

		eosScan, err := getEosAPI()
		if err != nil {
			fmt.Printf("error eos api: " + err.Error())
		}
		i := getPeriod(eosScan)
		if i == -1 {
			fmt.Printf("Period: Fail Period\n")
		}

		numToken := getNumberToken(1.0, eosScan[i].CreateOnDay, eosScan[i].DailyTotal)
		fmt.Printf("numToken per eth: \t\t\t%v\n", numToken)
		mustBuy(numToken, eosPriceUs, ethPriceUs)
		fmt.Printf("\n\n")
		time.Sleep(7 * time.Second)
	}
}

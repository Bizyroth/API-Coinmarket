// TODO create a package for all core functionality
// TODO create a cmd/coinmarket directory for the main.go file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// TODO use this endpoint to retrieve top 10 cryptomoney in Euro
// https://api.coinmarketcap.com/v1/ticker/?limit=10&convert=EUR

type api struct {
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

// Get specificat cryptomoney informations from coinmarketcap.com
func getAPICoinmarket(money string) ([]api, error) {
	var a []api
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

// TODO documentation
func getPeriod(ap []eosAPI) int {
	t := time.Now().UTC()
	for i := 0; i < 365; i++ {
		tB, _ := time.Parse(time.RFC3339, ap[i].Begins)
		tE, _ := time.Parse(time.RFC3339, ap[i].Ends)
		if tB.Before(t) && t.Before(tE) {
			return i
		}
	}
	return -1
}

// TODO: add documentation
func getNumberToken(totalOwnEth float64, totalEosToken float64, totalAllEth float64) float64 {
	return totalOwnEth * (totalEosToken / totalAllEth)
}

// TODO: add documentation
// return price Eos to buy, price Eso to sell, and percentage of gain if > 0.
func mustBuy(numToken float64, priceEos float64, priceEth float64) (float64, float64, float64) {
	priceEosInEth := 1 / numToken
	priceEosSell := priceEosInEth * priceEth
	percentGain := float64(0)

	if priceEos > priceEosSell {
		percentGain = 100 * priceEos / priceEosSell
	}
	return priceEosSell, priceEos, percentGain
}

var flagTime = flag.Int("time", 3, "update every X second")

func main() {
	flag.Parse()

	// TODO better naming
	fmt.Printf("\n%20s   %20s   %20s   %20s   %20s\n", "Price EOS in USD", "numToken per eth", "Price Eos to buy", "Price Eos to sell", "Percentage of gain")
	fmt.Printf("%20s   %20s   %20s   %20s   %20s", "...", "...", "...", "...", "...")

	// update periodically
	for range time.Tick(time.Duration(*flagTime) * time.Second) {
		// get EOS datas
		eosAP, err := getAPICoinmarket("eos")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		eosPriceUs, err := strconv.ParseFloat(eosAP[0].PriceUsd, 64)
		if err != nil {
			fmt.Println("Error conversion: " + err.Error())
		}
		eosScan, err := getEosAPI()
		if err != nil {
			fmt.Printf("error eos api: " + err.Error())
		}
		i := getPeriod(eosScan)
		if i == -1 {
			fmt.Printf("Period: Fail Period\n")
		}

		// get ether datas
		ethAP, err := getAPICoinmarket("ethereum")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		ethPriceUs, err := strconv.ParseFloat(ethAP[0].PriceUsd, 64)
		if err != nil {
			fmt.Println("Error conversion: " + err.Error())
		}

		// ???
		numToken := getNumberToken(1.0, eosScan[i].CreateOnDay, eosScan[i].DailyTotal)
		eosToBuy, eosToSell, gain := mustBuy(numToken, eosPriceUs, ethPriceUs)

		// TODO: detect Unix or Windows to use cursor position
		fmt.Printf("\r%20v   %20v   %20v   %20v   %20v", eosPriceUs, numToken, eosToBuy, eosToSell, gain)
	}
}

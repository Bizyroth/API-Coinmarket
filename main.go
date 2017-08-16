package main

import (
	"encoding/json"
	"fmt"
	//"os"
	"net/http"
	"io/ioutil"
	"time"
)



type API struct{
	Id					string `json:"id"`
	Name				string `json:"name"`
	Symbol				string `json:"symbol"`
	Rank				string `json:"rank"`
	Price_usd 			string `json:"price_usd"`
	Price_btc			string `json:"price_btc"`
	Vol_usd				string `json:"24h_volume_usd"`
	Market_cap_usd		string `json:"market_cap_usd"`
	Available_supply	string `json:"available_supply"`
	Total_supply		string `json:"total_supply"`
	Percent_change_1h	string `json:"percent_change_1h"`
	Percent24h			string `json:"percent_change_24h"`
	Percent7d			string `json:"percent_change_7d"`
	Last_updated		string `json:"last_updated"`
	Price_eur			string `json:"price_eur"`
	Vol_eur		 		string `json: "24h_volume_eur"`
	Market_cap_eur		string `json:"market_cap_eur"`

}

func main(){
	//i := 0
	url:="https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=10"

 for {
	cli := http.Client{}

	
	req, err:=cli.Get(url)
	if err != nil {
		fmt.Println("newrequest "+err.Error())
	}
	defer req.Body.Close()
	

	body, err := ioutil.ReadAll(req.Body)
	var aa [] API
	json.Unmarshal(body,&aa)
	fmt.Println("                id      symbol       price(eur)         %24h              %7j")
	fmt.Println()
	for l:= range aa {
	
		//fmt.Println(string(body))
		fmt.Printf("%20v     %3v      %15v       %6v      %10v", aa[l].Id,aa[l].Symbol,aa[l].Price_eur, 	aa[l].Percent24h, aa[l].Percent7d)
		fmt.Println()
	}
	time.Sleep(5 * time.Minute)
}		

	
}

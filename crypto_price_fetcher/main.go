package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup 

func GetPrice(index int, crypto string){
	defer wg.Done()
	url := fmt.Sprintf("https://api.cryptapi.io/%s/info/?prices=1", crypto)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("access-control-allow-origin", "*")
	res, _ := http.DefaultClient.Do(req)

	var result map[string]interface{}

	err := json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		fmt.Println(crypto,"Error decoding response:", err)
	}
	if prices, ok := result["prices"].(map[string]interface{}); ok {
		if usd_price, ok := prices["USD"]; ok {
			fmt.Printf("[%d] %s: %s USD\n", index, strings.ToUpper(crypto), usd_price)
		}else{
			fmt.Printf("[%d] No USD price found for %s \n", index, crypto)
		}
	}
}


func main(){
	
	var tickers []string
	tickers = append(tickers, "btc", "doge", "eth", "bch", "ltc")
	fmt.Printf("Getting Prices For: %v \n", tickers)

	for i, ticker := range tickers {
		go GetPrice(i, ticker)
		wg.Add(1)
	}
	wg.Wait()
}
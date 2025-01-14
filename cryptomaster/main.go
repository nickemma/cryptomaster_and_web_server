package main

import (
	"fmt"

	"github.com/nickemma/cryptomaster_and_web_server/api"
)

func main() {
	done := make(chan bool)

	currencies := []string{"BTC", "BCH", "ETH", "SOL"}
	for _, value := range currencies {
		go func() {
			getCurrencyData(value)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for range currencies {
		<-done
	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}

package main

import (
	"fmt"
	"sync"

	"jeanruiz.dev/go/cryptopmaster/api"
)

func main() {
	currencies := []string{"BTC", "ETC", "CHIV"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %2.f", rate.Currency, rate.Price)
	}

}

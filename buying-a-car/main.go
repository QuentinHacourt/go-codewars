package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	counter := 0
	capital := 0
	for canBuy(startPriceOld, capital, startPriceNew) {

		capital += savingperMonth
		counter++
	}
	// you
	return [2]int{counter, capital}
}

func canBuy(priceOld, capital, priceNew int) bool {
	return priceOld+capital > priceNew
}

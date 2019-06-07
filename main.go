package main

import (
	"fmt"
	"log"
	"net/http"
	vm "simple-web/vendingmachine"
)

func main() {
	//vendingmachine
	//	- insert
	//	- totalbalance
	//	- buy
	v := vm.NewVendingMachine()
	http.HandleFunc("/insertcoins", func(w http.ResponseWriter, r *http.Request) {
		coin := r.FormValue("coin")
		v.InsertCoins(coin)
		fmt.Fprintf(w, "%v", v.ShowTotalBalance())
	})

	http.HandleFunc("/buydrink", func(w http.ResponseWriter, r *http.Request) {
		drink := r.FormValue("drink")
		v.InsertCoins(drink)
		fmt.Fprintf(w, "%v", v.BuyDrink(drink))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

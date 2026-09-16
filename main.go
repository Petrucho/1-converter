package main

import "fmt"

func main() {
	const USD_EUR = 1.1
	const USD_RUB = 100.0
	const EUR_RUB = USD_EUR * USD_RUB
	var global_fromCurrency, global_toCurrency string
	var global_sum float64

	global_sum, global_fromCurrency, global_toCurrency = getInput()
	calcValues(global_sum, global_fromCurrency, global_toCurrency)
}

func getInput() (return_sum float64, return_fromCurrency, return_toCurrency string) {

	fmt.Print("Введите исходящую валюту: ")
	_, err := fmt.Scan(&return_fromCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Введите сумму: ")
	_, err = fmt.Scan(&return_sum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Введите целевую валюту: ")
	_, err = fmt.Scan(&return_toCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

func calcValues(sum_p float64, curr1_p string, curr2_p string) (return_convertedSum float64) {
	return
}

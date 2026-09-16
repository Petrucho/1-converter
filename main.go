package main

import "fmt"

func main() {
	const USD_EUR = 1.1
	const USD_RUB = 100.0
	const EUR_RUB = USD_EUR * USD_RUB
	var fromCurrency, toCurrency string
	var sum float64

	getInput()
	calcValues(sum, fromCurrency, toCurrency)
}

func getInput() {

	fmt.Print("Введите исходящую валюту: ")
	_, err := fmt.Scan(&fromCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Введите сумму: ")
	_, err := fmt.Scan(&sum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Введите целевую валюту: ")
	_, err := fmt.Scan(&toCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

func calcValues(sum_p float64, curr1_p string, curr2_p string) float64 {
	var convertedSum float64
	return convertedSum
}

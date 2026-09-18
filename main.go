package main

import (
	"fmt"
	"strings"
)

const USD_EUR = 1.1
const USD_RUB = 100.0
const EUR_RUB = USD_EUR * USD_RUB

func main() {
	fmt.Printf("Конвертированная сумма: %0.2f\n", calcValues(getInput()))
}

func getInput() (return_amount float64, return_fromCurrency, return_toCurrency string) {
	return_fromCurrency = getFromCurrency()
	return_amount = getAmount()
	return_toCurrency = getToCurrency(return_fromCurrency)
	return
}

func getFromCurrency() (return_fromCurrency string) {
outerLoop:
	for {
		fmt.Print("Введите исходящую валюту (USD, EUR, RUB): ")
		_, err := fmt.Scan(&return_fromCurrency)
		if err != nil {
			continue
		}
		return_fromCurrency = strings.ToUpper(return_fromCurrency)
		switch return_fromCurrency {
		case "USD", "EUR", "RUB":
			break outerLoop
		default:
			continue
		}
	}
	return
}

func getAmount() (return_amount float64) {
outerLoop:
	for {
		fmt.Print("Введите сумму больше нуля: ")
		_, err := fmt.Scan(&return_amount)
		if err != nil {
			continue
		}
		switch {
		case return_amount > 0:
			break outerLoop
		default:
			continue
		}
	}
	return
}

func getToCurrency(param_fromCurrency string) (return_toCurrency string) {
outerLoop:
	for {
		switch param_fromCurrency {
		case "USD":
			fmt.Print("Введите целевую валюту 'EUR', 'RUB': ")
		case "EUR":
			fmt.Print("Введите целевую валюту 'USD', 'RUB': ")
		case "RUB":
			fmt.Print("Введите целевую валюту 'USD', 'EUR': ")
		default:
			error_string := fmt.Sprintf("Должно остаться только USD или EUR или RUB!\nа оказалось %s", param_fromCurrency)
			panic(error_string)
		}
		_, err := fmt.Scan(&return_toCurrency)
		if err != nil {
			continue
		}
		return_toCurrency = strings.ToUpper(return_toCurrency)
		switch {
		case param_fromCurrency == return_toCurrency:
			continue
		case return_toCurrency == "USD" || return_toCurrency == "EUR" || return_toCurrency == "RUB":
			break outerLoop
		default:
			continue
		}
	}
	return
}

func calcValues(amount_p float64, fromCurrency_p string, toCurrency_p string) (return_convertedAmount float64) {
	//fmt.Printf("running calcValues with params:\namount_p: %0.2f\nfromCurrency_p: %s\ntoCurrency_p: %s\n", amount_p, fromCurrency_p, toCurrency_p)
	if (amount_p != 0) && (fromCurrency_p != "") && (toCurrency_p != "") {
		switch {
		case fromCurrency_p == "USD" && toCurrency_p == "EUR":
			return_convertedAmount = USD_EUR * amount_p
		case fromCurrency_p == "USD" && toCurrency_p == "RUB":
			return_convertedAmount = USD_RUB * amount_p
		case fromCurrency_p == "EUR" && toCurrency_p == "USD":
			return_convertedAmount = amount_p / USD_EUR
		case fromCurrency_p == "EUR" && toCurrency_p == "RUB":
			return_convertedAmount = EUR_RUB * amount_p
		case fromCurrency_p == "RUB" && toCurrency_p == "USD":
			return_convertedAmount = amount_p / USD_RUB
		case fromCurrency_p == "RUB" && toCurrency_p == "EUR":
			return_convertedAmount = amount_p / EUR_RUB
		}
	} else {
		fmt.Printf("Не задан один из параметров:\nСумма: %0.2f\nВалюта источник: %s\nВалюта целевая: %s\n", amount_p, fromCurrency_p, toCurrency_p)
		return_convertedAmount = 0
	}
	return
}

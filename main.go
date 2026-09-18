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

func getInput() (return_sum float64, return_fromCurrency, return_toCurrency string) {
outerLoop1:
	for {
		fmt.Print("Введите исходящую валюту (USD, EUR, RUB): ")
		fmt.Scan(&return_fromCurrency)

		switch strings.ToUpper(return_fromCurrency) {
		case "USD", "EUR", "RUB":
			break outerLoop1
		default:
			continue
		}
	}

outerLoop2:
	for {
		fmt.Print("Введите сумму больше нуля: ")
		_, err := fmt.Scan(&return_sum)
		switch {
		case return_sum > 0:
			break outerLoop2
		case err != nil:
			fmt.Printf("какая-то ошибка ввода суммы\n")
			continue
		default:
			continue
		}
	}

outerLoop3:
	for {
		switch strings.ToUpper(return_fromCurrency) {
		case "USD":
			fmt.Print("Введите целевую валюту 'EUR', 'RUB': ")
		case "EUR":
			fmt.Print("Введите целевую валюту 'USD', 'RUB': ")
		case "RUB":
			fmt.Print("Введите целевую валюту 'USD', 'EUR': ")
		default:
			error_string := fmt.Sprintf("Должно остаться только USD или EUR или RUB!\nа оказалось %s", return_fromCurrency)
			panic(error_string)
		}
		fmt.Scan(&return_toCurrency)
		switch strings.ToUpper(return_toCurrency) {
		case "USD", "EUR", "RUB":
			break outerLoop3
		default:
			continue
		}
	}

	return
}

func calcValues(amount_p float64, fromCurrency_p string, toCurrency_p string) (return_convertedAmount float64) {
	//fmt.Printf("running calcValues with params:\namount_p: %0.2f\nfromCurrency_p: %s\ntoCurrency_p: %s\n", amount_p, fromCurrency_p, toCurrency_p)
	if (amount_p != 0) || (fromCurrency_p != "") || (toCurrency_p != "") {
		switch {
		case strings.ToUpper(fromCurrency_p) == "USD" && strings.ToUpper(toCurrency_p) == "EUR":
			return_convertedAmount = USD_EUR * amount_p
		case strings.ToUpper(fromCurrency_p) == "USD" && strings.ToUpper(toCurrency_p) == "RUB":
			return_convertedAmount = USD_RUB * amount_p
		case strings.ToUpper(fromCurrency_p) == "EUR" && strings.ToUpper(toCurrency_p) == "USD":
			return_convertedAmount = amount_p - (USD_EUR / amount_p)
		case strings.ToUpper(fromCurrency_p) == "EUR" && strings.ToUpper(toCurrency_p) == "RUB":
			return_convertedAmount = EUR_RUB * amount_p
		case strings.ToUpper(fromCurrency_p) == "RUB" && strings.ToUpper(toCurrency_p) == "USD":
			return_convertedAmount = amount_p / USD_RUB
		case strings.ToUpper(fromCurrency_p) == "RUB" && strings.ToUpper(toCurrency_p) == "EUR":
			return_convertedAmount = amount_p / EUR_RUB
		}
	} else {
		fmt.Printf("Не задан один из параметров:\nСумма: %0.2f\nВалюта источник: %s\nВалюта целевая: %s\n", amount_p, fromCurrency_p, toCurrency_p)
		return_convertedAmount = 0
	}
	return
}

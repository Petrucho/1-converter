package main

import (
	"fmt"
	"strings"
)

type CurrencyMapType = map[string]map[string]float64

func main() {
	CurrencyMap := CurrencyMapType{"USD": {"EUR": 1.1, "RUB": 100.}, "EUR": {"USD": 1 / 1.1, "RUB": 110.}, "RUB": {"USD": 1 / 100., "EUR": 1 / 110.}}
outerLoop:
	for {
		switch getMenu() {
		case 1:
			returned_amount, returned_fromCurrency, returned_toCurrency := getInput()
			fmt.Printf("Конвертированная сумма: %0.2f\n", calcValues(CurrencyMap, returned_amount, returned_fromCurrency, returned_toCurrency))
		case 2:
			break outerLoop
		default:
			continue
		}
	}
}

func getMenu() (returnInput int) {
	for {
		fmt.Print("1. Конвертировать\n2. Выход\n")
		_, err := fmt.Scanf("%d", &returnInput)
		if err != nil {
			continue
		}
		break
	}
	return
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

func calcValues(currencyMap_p CurrencyMapType, amount_p float64, fromCurrency_p string, toCurrency_p string) (return_convertedAmount float64) {
	//fmt.Printf("running calcValues with params:\namount_p: %0.2f\nfromCurrency_p: %s\ntoCurrency_p: %s\n", amount_p, fromCurrency_p, toCurrency_p)
	if (amount_p != 0) && (fromCurrency_p != "") && (toCurrency_p != "") {
		returned_rate, found_rate := GetRate(currencyMap_p, fromCurrency_p, toCurrency_p)
		if found_rate {
			return_convertedAmount = returned_rate * amount_p
		} else {
			fmt.Printf("Не найден курс для %s→%s", fromCurrency_p, toCurrency_p)
		}
	} else {
		fmt.Printf("Не задан один из параметров:\nСумма: %0.2f\nВалюта источник: %s\nВалюта целевая: %s\n", amount_p, fromCurrency_p, toCurrency_p)
		return_convertedAmount = 0
	}
	return
}

func GetRate(rates CurrencyMapType, from, to string) (float64, bool) {
	inner, ok := rates[from]
	if !ok {
		return 0, false // нет такой исходной валюты
	}

	rate, ok := inner[to]
	return rate, ok
}

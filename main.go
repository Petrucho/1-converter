package main

import (
	"fmt"
	"strings"
)

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
	var endProgram bool = false

outerLoop1:
	for {
		fmt.Print("Введите исходящую валюту (USD, EUR, RUB)\nили для завершения программы нажмите Enter: ")
		fmt.Scan(&return_fromCurrency)

		switch strings.ToUpper(return_fromCurrency) {
		case "":
			fmt.Println("Выбрано завершение программы")
			endProgram = true
			break outerLoop1
		case "USD", "EUR", "RUB":
			break outerLoop1
		default:
			continue
		}
	}

	if !endProgram {
	outerLoop2:
		for {
			fmt.Print("Введите сумму больше нуля\nили для выхода из программы нажмите Enter: ")
			_, err := fmt.Scan(&return_sum)
			switch {
			case return_sum > 0:
				break outerLoop2
			case err != nil:
				endProgram = true
				break outerLoop2
			default:
				continue
			}
		}

		if !endProgram {
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
				case "":
					fmt.Println("Выбрано завершение программы")
					endProgram = true
					break outerLoop3
				case "USD", "EUR", "RUB":
					break outerLoop3
				default:
					continue
				}

			}
		}
	}
	return
}

func calcValues(sum_p float64, curr1_p string, curr2_p string) (return_convertedSum float64) {
	return
}

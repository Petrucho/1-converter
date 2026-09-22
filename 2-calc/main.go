package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var someSlice []int

	operationStr := getOperation()
	someSlice = getSlice()
	sort.Ints(someSlice)
	if len(someSlice) > 0 { // non-empty slice
		switch operationStr {
		case "AVG":
			fmt.Printf("AVG = %d\n", operationAVG(someSlice))
		case "SUM":
			fmt.Printf("SUM = %d\n", operationSUM(someSlice))
		case "MED":
			fmt.Printf("MED = %d\n", operationMED(someSlice))
		default:
			break
		}
	}
}

func getOperation() (return_Operation string) {
outerLoop:
	for {
		fmt.Print("Введите операцию (AVG, SUM, MED): ")
		_, err := fmt.Scan(&return_Operation)
		if err != nil {
			continue
		}
		return_Operation = strings.ToUpper(return_Operation)
		switch return_Operation {
		case "AVG", "SUM", "MED":
			break outerLoop
		default:
			continue
		}
	}
	return
}

func getSlice() (return_slice []int) {
	fmt.Print("Введите целые числа через запятую:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		// Разбиваем строку по запятым
		parts := strings.Split(strings.TrimSpace(line), ",")
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				return_slice = append(return_slice, num)
			} else {
				fmt.Printf("Ошибка преобразования '%s'\n", part)
			}
		}
	}
	return
}

func operationAVG(param_slice []int) int {
	return operationSUM(param_slice) / len(param_slice)
}

func operationSUM(param_slice []int) (return_SUM int) {
	for _, someValue := range param_slice {
		return_SUM += someValue
	}
	return
}

func operationMED(param_slice []int) (return_MED int) {
	numberCount := len(param_slice)
	if numberCount%2 == 0 {
		return_MED = (param_slice[numberCount/2] + param_slice[(numberCount/2)+1]) / 2
	} else {
		return_MED = (param_slice[(numberCount/2)+1])
	}
	return
}

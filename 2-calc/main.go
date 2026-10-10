package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var menuFunction = map[string]func([]int) int{
	"AVG": operationAVG,
	"SUM": operationSUM,
	"MED": operationMED,
}

func main() {
	var someSlice []int

	operationStr := getOperation()
	someSlice = getSlice()
	sort.Ints(someSlice)

	menuFunc := menuFunction[operationStr]
	if menuFunc == nil {
		fmt.Println("Не настроено меню функций!")
		return
	}

	if len(someSlice) > 0 {
		fmt.Printf("%s = %d\n", operationStr, menuFunc(someSlice)) // non-empty slice
		/*switch operationStr {
		case "AVG":
			fmt.Printf("AVG = %d\n", menuFunc(someSlice))
		case "SUM":
			fmt.Printf("SUM = %d\n", menuFunc(someSlice))
		case "MED":
			fmt.Printf("MED = %d\n", menuFunc(someSlice))
		default:
			break
			}*/
	} else {
		fmt.Printf("Slice is empty!")
	}
}

func getOperation() (return_Operation string) {
	keys := getKeys(menuFunction)
	sort.Strings(keys)
outerLoop:
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Введите операцию %s: ", keys)
		read_Operation, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		return_Operation = strings.ToUpper(strings.TrimSpace(read_Operation))
		if _, ok := menuFunction[return_Operation]; ok {
			break outerLoop
		}
		/*switch return_Operation {
		case "AVG", "SUM", "MED":
			break outerLoop
		default:
			continue
			}*/
	}
	return
}

func getSlice() (return_slice []int) {
	fmt.Print("Введите целые числа через запятую:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		// Разбиваем строку по запятым
		parts := strings.Split(line, ",")
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
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
		return_MED = (param_slice[(numberCount/2)-1] + param_slice[(numberCount/2)]) / 2
	} else {
		return_MED = (param_slice[(numberCount / 2)])
	}
	return
}

func getKeys(m map[string]func([]int) int) []string {
	keys := make([]string, 0, len(m)) // заранее резервируем ёмкость
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"fmt"
	"time"
)

func main() {
	myList := bins.NewBinList()

	// Создаем новый элемент (указатель на Bin)
	newBin, err := bins.CreateBin("bin-99", false, time.Now(), "My First Bin")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вызываем метод Add. Так как myList — это уже указатель,
	// Go сам корректно применит метод и обновит слайс внутри myList.
	myList.Add(newBin)

	// Проверяем результат пакета bins
	fmt.Printf("В списке элементов: %d\n", len(myList.Bins))
	fmt.Printf("Имя первого элемента: %s\n", myList.Bins[0].Name)

	// Проверяем результат пакета api
	api.ReturnSomeText()

	// Проверяем результат пакета file
	file.ReadFile()
	file.WriteFile()

	// Проверяем результат пакета storage
	storage.PrintSomething()
}

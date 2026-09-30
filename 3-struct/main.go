package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct{ Bins []Bin }

func (bl *BinList) Add(newBin *Bin) {
	if newBin != nil && bl != nil {
		bl.Bins = append(bl.Bins, *newBin)
	}
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}

func main() {
	myList := NewBinList()

	// Создаем новый элемент (указатель на Bin)
	newBin := &Bin{
		id:        "bin-99",
		name:      "My First Bin",
		createdAt: time.Now(),
	}

	// Вызываем метод Add. Так как myList — это уже указатель,
	// Go сам корректно применит метод и обновит слайс внутри myList.
	myList.Add(newBin)

	// Проверяем результат
	fmt.Printf("В списке элементов: %d\n", len(myList.Bins))
	fmt.Printf("Имя первого элемента: %s\n", myList.Bins[0].name)
}

func createBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	if createdAt.IsZero() {
		return nil, errors.New("INVALID_CREATEDAT")
	}
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
	return newBin, nil
}

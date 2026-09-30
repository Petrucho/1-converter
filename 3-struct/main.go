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

func main() {
	BinList := make([]Bin, 0)
	newBin := &Bin{
		id:   "1",
		name: "First Bin",
	}

	BinList = addBinToList(BinList, newBin)
	fmt.Println("Длина списка:", len(BinList))
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

func addBinToList(binList []Bin, newBin *Bin) []Bin {
	if newBin != nil {
		binList = append(binList, *newBin)
	}
	return binList
}

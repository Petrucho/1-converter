package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
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

func CreateBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
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
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
	return newBin, nil
}

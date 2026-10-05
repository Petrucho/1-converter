package bins

import (
	"errors"
	"fmt"
	"time"
)

var (
	// ErrInvalidID возвращается, если id пустой.
	ErrInvalidID = errors.New("INVALID_ID")
	// ErrInvalidCreatedAt возвращается, если время создания нулевое.
	ErrInvalidCreatedAt = errors.New("INVALID_CREATEDAT")
	// ErrInvalidName возвращается, если имя пустое.
	ErrInvalidName = errors.New("INVALID_NAME")
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

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
		return nil, ErrInvalidID
	}
	if createdAt.IsZero() {
		return nil, ErrInvalidCreatedAt
	}
	if name == "" {
		return nil, ErrInvalidName
	}
	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
	return newBin, nil
}

// Repository — то, куда сервис сохраняет список корзин и откуда читает его.
// Интерфейс объявлен рядом с потребителем: пакету bins не нужно знать, что
// реализация — это storage поверх файла, и не приходится его импортировать
// (иначе получился бы цикл: storage уже импортирует bins).
type Repository interface {
	SaveBins(list *BinList) error
	ReadBins() (*BinList, error)
}

// Service — операции над корзинами. Хранилище приходит снаружи
// (constructor injection), поэтому сервис не привязан ни к файлу, ни к формату.
type Service struct {
	repo Repository
	list *BinList
}

// NewService собирает сервис поверх хранилища.
func NewService(repo Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("bins: repository is nil")
	}
	return &Service{repo: repo, list: NewBinList()}, nil
}

// Create создаёт корзину и добавляет её в текущий список.
func (s *Service) Create(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if s == nil {
		return nil, errors.New("bins: service is nil")
	}
	b, err := CreateBin(id, private, createdAt, name)
	if err != nil {
		return nil, err
	}
	s.List().Add(b)
	return b, nil
}

// List возвращает текущий список корзин (никогда не nil).
func (s *Service) List() *BinList {
	if s.list == nil {
		s.list = NewBinList()
	}
	return s.list
}

// Save сохраняет текущий список через хранилище.
func (s *Service) Save() error {
	if s == nil {
		return errors.New("bins: service is nil")
	}
	if err := s.repo.SaveBins(s.List()); err != nil {
		return fmt.Errorf("bins: save: %w", err)
	}
	return nil
}

// Load читает список из хранилища и запоминает его как текущий.
func (s *Service) Load() (*BinList, error) {
	if s == nil {
		return nil, errors.New("bins: service is nil")
	}
	list, err := s.repo.ReadBins()
	if err != nil {
		return nil, fmt.Errorf("bins: load: %w", err)
	}
	if list == nil {
		list = NewBinList()
	}
	s.list = list
	return list, nil
}

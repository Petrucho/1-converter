package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"3-struct/bins"
)

// DefaultFileName — файл, в который сохраняются корзины по умолчанию.
const DefaultFileName = "bins_storage.json"

// FileStore — то, что Storage нужно от файлового хранилища.
// Интерфейс объявлен рядом с потребителем и содержит только используемые
// методы, поэтому storage не импортирует пакет file: подойдёт *file.JsonDb
// или любой мок в памяти.
type FileStore interface {
	// ReadFile читает файл целиком. Если файла нет, ошибка оборачивает os.ErrNotExist.
	ReadFile() ([]byte, error)
	// WriteFile перезаписывает файл содержимым content.
	WriteFile(content []byte) error
}

// Storage сериализует список корзин в JSON и читает его обратно.
// Файловое хранилище приходит снаружи (dependency injection): сам Storage
// не создаёт его и не знает, что там за реализация — только интерфейс FileStore.
type Storage struct {
	repo FileStore
}

// New собирает Storage поверх файлового хранилища.
func New(repo FileStore) (*Storage, error) {
	if repo == nil {
		return nil, errors.New("storage: repo is nil")
	}
	return &Storage{repo: repo}, nil
}

// SaveBins сериализует список в JSON и сохраняет в файл.
func (s *Storage) SaveBins(list *bins.BinList) error {
	if list == nil {
		return errors.New("bin list is nil")
	}

	// Маршалим структуру BinList в JSON (с отступами для читаемости)
	jsonData, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal bins: %w", err)
	}

	if err := s.repo.WriteFile(jsonData); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	return nil
}

// ReadBins читает файл и десериализует его обратно в структуру BinList.
// Если файла ещё нет — возвращает пустой список, а не ошибку.
func (s *Storage) ReadBins() (*bins.BinList, error) {
	jsonData, err := s.repo.ReadFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return bins.NewBinList(), nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Создаем новый пустой список с помощью конструктора
	list := bins.NewBinList()
	if len(jsonData) == 0 {
		return list, nil
	}

	// Наполняем структуру данными из JSON
	if err := json.Unmarshal(jsonData, list); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bins: %w", err)
	}
	return list, nil
}

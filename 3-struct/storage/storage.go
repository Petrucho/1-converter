package storage

import (
	"encoding/json"
	"fmt"

	"3-struct/bins"
	"3-struct/file"
)

// Имя файла, в который будем сохранять корзины
const storageFileName = "bins_storage.json"

func PrintSomething() {
	fmt.Println("Print from storage package")
}

// SaveBins сериализует список в JSON и сохраняет в файл
func SaveBins(list *bins.BinList) error {
	if list == nil {
		return fmt.Errorf("bin list is nil")
	}

	// Маршалим структуру BinList в JSON (с отступами для читаемости)
	jsonData, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal bins: %w", err)
	}

	// Используем пакет file для записи
	err = file.WriteFile(jsonData, storageFileName)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

// ReadBins читает файл и десериализует его обратно в структуру BinList
func ReadBins() (*bins.BinList, error) {
	// Используем пакет file для чтения
	jsonData, err := file.ReadFile(storageFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Создаем новый пустой список с помощью конструктора
	list := bins.NewBinList()

	// Наполняем структуру данными из JSON
	err = json.Unmarshal(jsonData, list)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bins: %w", err)
	}

	return list, nil
}

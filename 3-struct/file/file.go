package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if fileExists(name) {
		// Используем os.ReadFile для чтения всего файла в память
		data, err := os.ReadFile(name)
		if err != nil {
			// Если произошла ошибка, возвращаем nil вместо данных и саму ошибку
			return nil, err
		}
		// Если всё прошло успешно, возвращаем данные и nil вместо ошибки
		return data, nil
	} else {
		fmt.Printf("Файл не найден или не разрешён доступ: %s\n", name)
	}
	return nil, nil
}

func WriteFile(content []byte, name string) error {
	fmt.Println("Write file:", name)
	// os.WriteFile записывает данные в файл.
	// 0664 — это права доступа (чтение/запись для владельца и группы, чтение для остальных)
	err := os.WriteFile(name, content, 0664)
	if err != nil {
		// Если возникла ошибка (например, нет прав или папки), возвращаем её
		return err
	}

	// Если всё прошло успешно, возвращаем nil
	return nil
}
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // File exists
	}
	if errors.Is(err, os.ErrNotExist) {
		return false // File explicitly does not exist
	}
	// The file might exist, but we got a different error (e.g., permission denied)
	return false
}

func IsJSON(name string) bool {
	if strings.ToLower(filepath.Ext(name)) == ".json" {
		return true
	}
	return false
}

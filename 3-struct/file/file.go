package file

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	fmt.Println("Read file:", name)
	// Используем os.ReadFile для чтения всего файла в память
	data, err := os.ReadFile(name)
	if err != nil {
		// Если произошла ошибка, возвращаем nil вместо данных и саму ошибку
		return nil, err
	}

	// Если всё прошло успешно, возвращаем данные и nil вместо ошибки
	return data, nil
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

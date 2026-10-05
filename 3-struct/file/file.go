package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Repo — хранилище, работающее с одним файлом.
// Потребители (storage и т.д.) зависят от этого интерфейса, а не от *JsonDb,
// поэтому реализацию легко подменить: в тестах — мок в памяти, в проде — диск.
type Repo interface {
	// Name возвращает имя файла, с которым работает реализация.
	Name() string
	// ReadFile читает файл целиком. Если файла нет, ошибка оборачивает os.ErrNotExist.
	ReadFile() ([]byte, error)
	// WriteFile перезаписывает файл содержимым content.
	WriteFile(content []byte) error
}

// JsonDb — реализация Repo поверх локальной файловой системы.
type JsonDb struct {
	filename string
}

// Проверка на этапе компиляции: JsonDb удовлетворяет интерфейсу Repo.
var _ Repo = (*JsonDb)(nil)

// NewJsonDb создаёт репозиторий для конкретного файла.
func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

// Name возвращает имя файла.
func (db *JsonDb) Name() string {
	return db.filename
}

// ReadFile читает весь файл в память.
func (db *JsonDb) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", db.filename, err)
	}
	return data, nil
}

// WriteFile записывает данные в файл.
// 0664 — это права доступа (чтение/запись для владельца и группы, чтение для остальных).
func (db *JsonDb) WriteFile(content []byte) error {
	if err := os.WriteFile(db.filename, content, 0664); err != nil {
		return fmt.Errorf("write %q: %w", db.filename, err)
	}
	return nil
}

// IsJSON проверяет, что у файла расширение .json.
func IsJSON(name string) bool {
	return strings.EqualFold(filepath.Ext(name), ".json")
}

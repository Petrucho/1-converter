package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// JsonDb — JSON-файл на локальной файловой системе.
// Пакет file не объявляет интерфейсов: он просто возвращает конкретный тип,
// а потребители (storage и т.д.) сами описывают нужные им методы.
type JsonDb struct {
	filename string
}

// NewJsonDb создаёт хранилище для конкретного файла.
// Возвращает ошибку, если у файла не расширение .json.
func NewJsonDb(name string) (*JsonDb, error) {
	if !IsJSON(name) {
		return nil, fmt.Errorf("file: %q is not a JSON file", name)
	}
	return &JsonDb{
		filename: name,
	}, nil
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

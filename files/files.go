// Package files for read and write files
package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// Тип имеющий методы ReadFile и WriteFile, интерфейс в (storage package) имеет контракт с этим типом
type JSONDB struct {
	fileName string
}

// Добавляем название файлам
func NewJSONDB(name string) *JSONDB {
	return &JSONDB{
		fileName: name,
	}
}

// Чтение файла из локального компьютера
func (db *JSONDB) ReadFile() ([]byte, error) {
	_, err := os.Stat(db.fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Файла не существует")
	} else if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла")
	}
	ext := filepath.Ext(db.fileName)
	if ext != ".json" {
		return nil, fmt.Errorf("Файл не имеет расширения json")
	}
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла")
	}
	return data, nil
}

// Запись в файл
func (db *JSONDB) WriteFile(data []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		return fmt.Errorf("Ошибка создании файла")
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл")
	}
	return nil
}

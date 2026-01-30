// Package files for read and write files
package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Файла не существует")
	} else if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла")
	}
	ext := filepath.Ext(fileName)
	if ext != ".json" {
		return nil, fmt.Errorf("Файл не имеет расширения json")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла")
	}
	return data, nil
}

func WriteFile(data []byte, fileName string) error {
	file, err := os.Create(fileName)
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

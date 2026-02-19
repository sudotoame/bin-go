// Package storage for storage in file
package storage

import (
	"dz/bingo/bins"
	"encoding/json"
	"fmt"
	"time"
)

// Имеет контракты с методами JSONDB(files package)
type DB interface {
	ReadFile() ([]byte, error)
	WriteFile([]byte) error
}

// Слайс баз данных Bin
type Vault struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
}

// Добавление зависимости
type VaultWithDB struct {
	Vault
	Db DB
}

// Method ToByte - работает как метод Vault из за прямой нужды работы с базой данных а не со слайсом этих данных
func (vault *Vault) ToByte() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, fmt.Errorf("Ошибка сериализации")
	}
	return data, nil
}

// func save - обновляет время обновления файл указанный в константе main пакета.
// Сериализует данные из структуры и добавляет данные в файл
func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = vault.Db.WriteFile(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Добавляет новые элементы в Vault, добавляя новые записи в файл
func (vault *VaultWithDB) AddBin(binAdd bins.Bin) {
	vault.Bins = append(vault.Bins, binAdd)
	vault.save()
}

// Десериализует прочитанный файл с помощью метода ReadFile и добавляет в Vault.
// Обновляет наш слайс базы данных.
// Работает через контракты DB
func NewVault(DB DB) *VaultWithDB {
	file, err := DB.ReadFile()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Bins:      []bins.Bin{},
				UpdatedAt: time.Now(),
			},
			Db: DB,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Bins:      []bins.Bin{},
				UpdatedAt: time.Now(),
			},
			Db: DB,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		Db:    DB,
	}
}

func (vault *VaultWithDB) DeleteBin(id string) bool {
	for i, v := range vault.Bins {
		if id == v.ID {
			vault.Bins = append(vault.Bins[:i], vault.Bins[i+1:]...)
			vault.save()
			return true
		}
	}
	return false
}

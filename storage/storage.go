// Package storage for storage in file
package storage

import (
	"dz/bingo/bins"
	"dz/bingo/files"
	"encoding/json"
	"fmt"
	"time"
)

type Vault struct {
	Bins      []bins.Bin
	UpdatedAt time.Time
}

func (vault *Vault) ToByte() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, fmt.Errorf("Ошибка сериализации")
	}
	return data, nil
}

func (vault *Vault) AddBin(binAdd bins.Bin) {
	vault.Bins = append(vault.Bins, binAdd)
	vault.UpdatedAt = time.Now()
}

func NewVault(fileName string) *Vault {
	file, err := files.ReadFile(fileName)
	if err != nil {
		return &Vault{
			Bins:      []bins.Bin{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &vault
}

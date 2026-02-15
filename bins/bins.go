// Package bins for bin struct
package bins

import (
	"fmt"
)

// Структура самой базы данных
type Bin struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Private   bool   `json:"private"`
	Name      string `json:"name"`
}

// Создание одной базы данных
func NewBin(id, name, time string, private bool) (*Bin, error) {
	if id == "" {
		return nil, fmt.Errorf("Invalid ID")
	}
	if name == "" {
		return nil, fmt.Errorf("Invalid name")
	}
	newBin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time,
		Name:      name,
	}
	return newBin, nil
}

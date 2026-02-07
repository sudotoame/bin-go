// Package bins for bin struct
package bins

import (
	"fmt"
	"time"
)

// Структура самой базы данных
type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

// Создание одной базы данных
func NewBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, fmt.Errorf("Invalid ID")
	}
	if name == "" {
		return nil, fmt.Errorf("Invalid name")
	}
	newBin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin, nil
}

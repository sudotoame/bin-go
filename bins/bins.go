// Package bins for bin struct
package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

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

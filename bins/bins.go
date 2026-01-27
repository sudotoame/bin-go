// Package bins for bin struct
package bins

import "time"

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id, name string, private bool, createdAt time.Time) (*Bin, error) {
	return &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

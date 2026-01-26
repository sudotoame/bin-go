package main

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func newBin(id, name string, private bool, createdAt time.Time) (*Bin, error) {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

func main() {
	var BinList []Bin
	BinList = append(BinList, Bin{
		Id:        "1",
		Private:   true,
		CreatedAt: time.Now(),
		Name:      "hi",
	})
	fmt.Println(BinList)
	res, err := newBin("2", "bye", false, time.Now())
	if err != nil {
		return
	}
	BinList = append(BinList, *res)
	fmt.Println(BinList)
}

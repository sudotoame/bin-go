package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id, name string, private bool, createdAt time.Time) (*Bin, error) {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}

func main() {
	var BinList []Bin
	BinList = append(BinList, Bin{
		id:        "1",
		private:   true,
		createdAt: time.Now(),
		name:      "hi",
	})
	fmt.Println(BinList)
	res, err := newBin("2", "bye", false, time.Now())
	if err != nil {
		return
	}
	BinList = append(BinList, *res)
	fmt.Println(BinList)
}

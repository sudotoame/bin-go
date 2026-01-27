package main

import (
	"dz/bingo/bins"
	"fmt"
	"time"
)

func main() {
	var BinList []bins.Bin
	BinList = append(BinList, bins.Bin{
		ID:        "1",
		Private:   true,
		CreatedAt: time.Now(),
		Name:      "hi",
	})
	fmt.Println(BinList)
	res, err := bins.NewBin("2", "bye", false, time.Now())
	if err != nil {
		return
	}
	BinList = append(BinList, *res)
	fmt.Println(BinList)
}

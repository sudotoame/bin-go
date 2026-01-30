package main

import (
	"dz/bingo/bins"
	"dz/bingo/files"
	"dz/bingo/storage"
	"fmt"
)

const fileName = "data.json"

func main() {
	newVault := storage.NewVault(fileName)

	createBin(newVault)

	fmt.Println(newVault.Bins)
	data, err := newVault.ToByte()
	err = files.WriteFile(data, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createBin(vault *storage.Vault) {
	privateCheck := false
	id := promtData("Введите ID")
	name := promtData("Введите name")
	private := promtData("Сделать приватной?(false default or press Y for true)")
	if private == "Y" || private == "y" {
		privateCheck = true
	}
	nBin, err := bins.NewBin(id, name, privateCheck)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddBin(*nBin)
}

func promtData(message string) string {
	var ch string
	fmt.Print(message + ": ")
	fmt.Scan(&ch)
	return ch
}

package main

import (
	"fmt"

	"dz/bingo/api"
	"dz/bingo/files"
	"dz/bingo/storage"

	"github.com/joho/godotenv"
)

const collectionFile = "bins.json"
const fileName = "data.json"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env load error")
	}
	newVault := storage.NewVault(files.NewJSONDB(collectionFile))
	createBin(newVault)
	// fmt.Println(newVault.Bins)

}

// Создание базы данных
func createBin(vault *storage.VaultWithDB) {
	privateCheck := false
	name := promtData([]string{"Введите name"})
	private := promtData([]string{"Сделать приватной?(false default or press Y for true)"})
	if private == "Y" || private == "y" {
		privateCheck = true
	}
	var privateData string
	if privateCheck == true {
		privateData = "true"
	} else {
		privateData = "false"
	}
	data, err := api.JsonBinPost(fileName, name, privateData)
	// nBin, err := bins.NewBin(name, privateCheck)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddBin(data.Metadata)
}

func promtData[T any](message []T) string {
	for i, v := range message {
		if i == len(message)-1 {
			fmt.Print(v, " :")
		} else {
			fmt.Println(v)
		}
	}
	var ch string
	if _, err := fmt.Scanln(&ch); err != nil {
		fmt.Println(err)
		return ""
	}
	return ch
}

package main

import (
	"flag"
	"fmt"

	"dz/bingo/api"
	"dz/bingo/files"
	"dz/bingo/storage"

	"github.com/joho/godotenv"
)

const collectionFile = "bins.json"

// TODO: Добавить флаги
// TODO: Добавить другие методы
// TODO: Добавить локального чтения айдишников и имен бинов
func main() {
	fileName := flag.String("file", "", "Название файла для создания бина")
	binName := flag.String("name", "", "Название бина")
	create := flag.Bool("create", false, "Создает бин если есть")
	help := flag.Bool("help", false, "Показывает какие флаги есть")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env load error")
	}
	newVault := storage.NewVault(files.NewJSONDB(collectionFile))
	if *create {
		createBin(newVault, *binName, *fileName)
		return
	}
	if *help {
		fmt.Println("флаг --create имеет два зависимых флага --file=filename файл который отправляем в jsonbin и --name==binname название бина")
		return
	}
	fmt.Println("Запущено без флагов: --help для помощи")
}

// Создание базы данных
func createBin(vault *storage.VaultWithDB, name string, fileName string) {
	privateCheck := false
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

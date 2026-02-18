package main

import (
	"flag"
	"fmt"
	"os"

	"dz/bingo/api"
	"dz/bingo/config"
	"dz/bingo/files"
	"dz/bingo/storage"

	"github.com/joho/godotenv"
)

const collectionFile = "bins.json"

// TODO: Добавить другие методы
// TODO: Добавить локальное чтения бинов по айдишнику и вывод имен бинов
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование:\n")
		fmt.Fprintf(os.Stderr, "  %s --create --file=data.json --name=my-bin [--private]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Флаги:")
		flag.PrintDefaults()
	}
	create := flag.Bool("create", false, "Создает бин если есть")
	fileName := flag.String("file", "", "Название файла для создания бина")
	binName := flag.String("name", "noname", "Название бина")
	private := flag.Bool("private", false, "Сделать бин приватным")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env load error")
		os.Exit(1)
	}
	key := config.NewConfig()
	if err = key.Validate(); err != nil {
		fmt.Println("Ошибка конфигурации: ", err)
		os.Exit(1)
	}
	apiNew := api.NewClient(key.Key)
	newVault := storage.NewVault(files.NewJSONDB(collectionFile))
	if *create {
		if *fileName == "" {
			fmt.Println("Файл должен быть задан")
			flag.Usage()
			os.Exit(2)
		}
		newJson := files.NewJSONDB(*fileName)
		myData, err := newJson.ReadFile()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		createBin(newVault, *binName, myData, *apiNew, *private)
		return
	}
}

// Создание базы данных
func createBin(vault *storage.VaultWithDB, binName string, myData []byte, newApi api.Client, privateData bool) {
	data, err := newApi.JsonBinPost(myData, binName, privateData)
	// nBin, err := bins.NewBin(name, privateCheck)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddBin(data.Metadata)
}

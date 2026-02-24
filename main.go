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

func main() {
	flagUsageInit()

	create := flag.Bool("create", false, "Создает бин если есть")
	get := flag.Bool("get", false, "Вывод бина по айдишнику")
	fileName := flag.String("file", "", "Название файла для создания бина")
	binName := flag.String("name", "noname", "Название бина")
	id := flag.String("id", "", "Айди бина")
	list := flag.Bool("list", false, "Прочитать локально сохраненные бины")
	update := flag.Bool("update", false, "Обновить бин по айдишнику")
	delete := flag.Bool("delete", false, "Удаляет бин по айди")
	private := flag.Bool("private", false, "Сделать бин приватным (Только для создания)")
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
		createBin(newVault, *fileName, *binName, *apiNew, *private)
		return
	}
	if *get {
		getBin(*id, *apiNew)
		return
	}
	if *list {
		listBin(*newVault)
		return
	}
	if *update {
		putBin(*fileName, *id, *apiNew)
		return
	}
	if *delete {
		deleteBin(*id, *apiNew, *newVault)
		return
	}
	flag.Usage()
}

func flagUsageInit() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование:\n")
		fmt.Fprintf(os.Stderr, "  %s --create --file=data.json --name=my-bin [--private]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --get --id=binid\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --delete --id=binid\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --update --file=data.json --id=binid\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --list\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Флаги:")
		flag.PrintDefaults()
	}
}

func getBin(id string, newApi api.Client) {
	if id == "" {
		fmt.Println("Айди не может быть пустым")
		flag.Usage()
		os.Exit(2)
	}
	data, err := newApi.GetBin(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data.Record))
}

func createBin(vault *storage.VaultWithDB, fileName string, binName string, newApi api.Client, privateData bool) {
	if fileName == "" {
		fmt.Println("Файл должен быть задан")
		flag.Usage()
		os.Exit(2)
	}
	newJson := files.NewJSONDB(fileName)
	myData, err := newJson.ReadFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	data, err := newApi.PostBin(myData, binName, privateData)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddBin(data.Metadata)
}

func listBin(vault storage.VaultWithDB) {
	data, err := vault.Db.ReadFile()
	if err != nil {
		fmt.Printf("Не удалось прочитать файл (list): %v", err)
		return
	}
	fmt.Println(string(data))
}

func putBin(fileName string, id string, newApi api.Client) {
	if fileName == "" {
		fmt.Println("Файл должен быть задан")
		flag.Usage()
		os.Exit(2)
	}
	if id == "" {
		fmt.Println("Айди не может быть пустым")
		flag.Usage()
		os.Exit(2)
	}
	newJson := files.NewJSONDB(fileName)
	myData, err := newJson.ReadFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	updData, err := newApi.UpdateBin(myData, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(updData.Record))
}

func deleteBin(id string, newApi api.Client, vaultBin storage.VaultWithDB) {
	if id == "" {
		fmt.Println("Айди не может быть пустым")
		flag.Usage()
		os.Exit(2)
	}
	if _, err := newApi.DeleteBin(id); err != nil {
		fmt.Println(err)
		return
	}
	if err := vaultBin.DeleteBin(id); !err {
		fmt.Println("Локальный бин не удален")
		return
	}
	fmt.Printf("Бин %s удален\n", id)
}

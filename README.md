#bin go

- `go run main.go` без флагов создает `data.json`
- `go run main.go --create --file="data.json" --name=my-bin` - Создает `bin` в JsonBin из `data.json`
- `go run main.go --update --file="data.json" --id=binid` - Обновляет `bin` в JsonBin из локального `data.json` 
- `go run main.go --delete --id=binid` - Удаляет бин по айдишнику
- `go run main.go --get --id=binid` - Выводит бин по айдишнику
- `go run main.go --list` 
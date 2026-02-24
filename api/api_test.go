package api_test

import (
	"dz/bingo/api"
	"dz/bingo/config"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func Init() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env load error")
		return "", fmt.Errorf("Ошибка загрузки .env файла: %v", err)
	}

	k := config.NewConfig()
	if err := k.Validate(); err != nil {
		return "", fmt.Errorf("Ошибка конфигурации: %v", err)
	}
	return k.Key, nil
}

type TestCases struct {
	Data    []byte
	Name    string
	Private bool
}

type MyData struct {
	Name  string
	Login string
}

func initData() *MyData {
	return &MyData{
		Name:  "Polka",
		Login: "Popo",
	}
}

func initTestCase() *TestCases {
	data := initData()
	myData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return &TestCases{
		Data:    myData,
		Name:    "Test",
		Private: false,
	}
}

func TestPostBin(t *testing.T) {
	key, err := Init()
	if err != nil {
		t.Skip(err)
	}
	if key == "" {
		t.Skip("Пропуск: Пустой API ключ")
	}
	client := api.NewClient(key)
	myTestCase := initTestCase()

	resp, err := client.PostBin(myTestCase.Data, myTestCase.Name, myTestCase.Private)
	if err != nil {
		t.Fatalf("PostBin: %v", err)
	}
	if resp == nil {
		t.Fatal("Ожидался non-nil response")
	}
	if resp.Metadata.ID == "" {
		t.Error("Metadata.ID Пустая")
	}
	if resp.Metadata.Name != myTestCase.Name {
		t.Errorf("Metadata.Name должен быть: %v, является: %v", myTestCase.Name, resp.Metadata.Name)
	}
	if resp.Metadata.Private != myTestCase.Private {
		t.Errorf("Metadata.Private должен быть: %v, является: %v", myTestCase.Private, resp.Metadata.Private)
	}
	if !strings.Contains(string(resp.Record), "Polka") || !strings.Contains(string(resp.Record), "Popo") {
		t.Errorf("Record должен содержать: %v, сейчас содержит: %v", myTestCase.Data, resp.Record)
	}
	defer client.DeleteBin(resp.Metadata.ID)
}

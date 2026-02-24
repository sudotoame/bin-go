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

func initData(name, login string) *MyData {
	return &MyData{
		Name:  name,
		Login: login,
	}
}

func initTestCase(data MyData) *TestCases {
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

func getAPIkey(t *testing.T) string {
	t.Helper()
	key, err := Init()
	if err != nil {
		t.Skip(err)
	}
	if key == "" {
		t.Skip("Пропуск: Пустой API ключ")
	}
	return key
}

func createTestBin(t *testing.T, myTestCase *TestCases, client *api.Client) *api.JsonBinResponse {
	t.Helper()
	resp, err := client.PostBin(myTestCase.Data, myTestCase.Name, myTestCase.Private)
	if err != nil {
		t.Fatalf("PostBin: %v", err)
	}
	if resp == nil {
		t.Fatal("Ожидался non-nil response (create bin)")
	}
	return resp
}

func assertValidBinResponse(t *testing.T, resp api.JsonBinResponse, myTestCase TestCases, recordExpected []string) {
	t.Helper()
	if resp.Metadata.ID == "" {
		t.Error("Metadata.ID Пустая")
	}
	if resp.Metadata.Name != myTestCase.Name {
		t.Errorf("Metadata.Name должен быть: %v, является: %v", myTestCase.Name, resp.Metadata.Name)
	}
	if resp.Metadata.Private != myTestCase.Private {
		t.Errorf("Metadata.Private должен быть: %v, является: %v", myTestCase.Private, resp.Metadata.Private)
	}
	if !strings.Contains(string(resp.Record), recordExpected[0]) || !strings.Contains(string(resp.Record), recordExpected[1]) {
		t.Errorf("Record должен содержать: %v, сейчас содержит: %v", myTestCase.Data, resp.Record)
	}
}
func TestPostBin(t *testing.T) {
	key := getAPIkey(t)
	recordExpected := []string{"Polka", "Popo"}

	client := api.NewClient(key)
	dataInit := initData(recordExpected[0], recordExpected[1])
	myTestCase := initTestCase(*dataInit)

	resp := createTestBin(t, myTestCase, client)
	defer client.DeleteBin(resp.Metadata.ID)
	assertValidBinResponse(t, *resp, *myTestCase, recordExpected)
}

func TestUpdateBin(t *testing.T) {
	key := getAPIkey(t)
	recordCreate := []string{"Polka", "Popo"}
	recordUpdated := []string{"Molka", "Momo"}

	client := api.NewClient(key)
	dataInit := initData(recordCreate[0], recordCreate[1])
	tc := initTestCase(*dataInit)

	resp := createTestBin(t, tc, client)
	defer client.DeleteBin(resp.Metadata.ID)

	updatedData := initData(recordUpdated[0], recordUpdated[1])
	updatedTc := initTestCase(*updatedData)

	updated, err := client.UpdateBin(updatedTc.Data, resp.Metadata.ID)
	if err != nil {
		t.Fatalf("UpdateBin: %v", err)
	}
	if updated == nil {
		t.Fatal("Ожидается non-nil response(update bin)")
	}
	assertValidBinResponse(t, *updated, *updatedTc, recordUpdated)
}

func TestGetBin(t *testing.T) {
	key := getAPIkey(t)
	recordExpected := []string{"Polka", "Popo"}

	client := api.NewClient(key)
	dataInit := initData(recordExpected[0], recordExpected[1])
	myTestCase := initTestCase(*dataInit)

	resp := createTestBin(t, myTestCase, client)
	defer client.DeleteBin(resp.Metadata.ID)
	getbin, err := client.GetBin(resp.Metadata.ID)
	if err != nil {
		t.Fatalf("GetBin: %v", err)
	}
	if getbin == nil {
		t.Fatal("Ожидается non-nil response(get bin)")
	}
	assertValidBinResponse(t, *getbin, *myTestCase, recordExpected)
}
func TestDeleteBin(t *testing.T) {
	key := getAPIkey(t)
	recordExpected := []string{"Polka", "Popo"}

	client := api.NewClient(key)
	dataInit := initData(recordExpected[0], recordExpected[1])
	myTestCase := initTestCase(*dataInit)

	resp := createTestBin(t, myTestCase, client)
	defer client.DeleteBin(resp.Metadata.ID)
	delbin, err := client.DeleteBin(resp.Metadata.ID)
	if err != nil {
		t.Fatalf("GetBin: %v", err)
	}
	if delbin == nil {
		t.Fatal("Ожидается non-nil response(get bin)")
	}
	assertValidBinResponse(t, *delbin, *myTestCase, recordExpected)
}

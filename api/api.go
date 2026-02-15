// Package api for work with json.bin
package api

import (
	"bytes"
	"dz/bingo/bins"
	"dz/bingo/config"
	"dz/bingo/files"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonBinResponse struct {
	// Record   json.RawMessage `json:"record"`
	Metadata bins.Bin `json:"metadata"`
}

const urlPost = "https://api.jsonbin.io/v3/b"

func JsonBinPost(filename string, binName string, private string) (*JsonBinResponse, error) {
	key := config.NewConfig()
	newJson := files.NewJSONDB(filename)
	myData, err := newJson.ReadFile()
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %v", err)
	}
	req, err := http.NewRequest("POST", urlPost, bytes.NewBuffer(myData))
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода POST: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key.Key) // Обязательно замените на реальный ключ
	// Опционально: можно задать имя бина
	req.Header.Set("X-Bin-Name", binName)
	req.Header.Set("X-Bin-Private", private)
	// Если нужно поместить бин в коллекцию, укажите X-Collection-Id
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения метода POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(req.Body)
		return nil, fmt.Errorf("Сервер вернул ошибку: %d, тело: %s", resp.StatusCode, body)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения ответа метода POST: %v", err)
	}

	var response JsonBinResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	fmt.Printf("Статус: %v\n", resp.Status)
	fmt.Printf("Ответ сервера: %v\n", string(body))
	return &response, nil
}

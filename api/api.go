// Package api for work with json.bin
package api

import (
	"bytes"
	"dz/bingo/bins"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type JsonBinResponse struct {
	// Record   json.RawMessage `json:"record"`
	Metadata bins.Bin `json:"metadata"`
}

type Client struct {
	MasterKey string
	HTTP      *http.Client
}

func NewClient(masterKey string) *Client {
	return &Client{
		MasterKey: masterKey,
		HTTP: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

const urlPost = "https://api.jsonbin.io/v3/b"

func (c *Client) JsonBinPost(myData []byte, binName string, private bool) (*JsonBinResponse, error) {
	req, err := http.NewRequest("POST", urlPost, bytes.NewBuffer(myData))
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода POST: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.MasterKey) // Обязательно замените на реальный ключ
	// Опционально: можно задать имя бина
	req.Header.Set("X-Bin-Name", binName)
	if private {
		req.Header.Set("X-Bin-Private", "true")
	} else {
		req.Header.Set("X-Bin-Private", "false")
	}
	// Если нужно поместить бин в коллекцию, укажите X-Collection-Id
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения метода POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
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

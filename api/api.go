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
	Record   json.RawMessage `json:"record"`
	Metadata bins.Bin        `json:"metadata"`
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

const baseUrl = "https://api.jsonbin.io/v3"

func (c *Client) DeleteBin(id string) error {
	url := fmt.Sprintf("%s/b/%s", baseUrl, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("Ошибка создани метода DELETE: %w", err)
	}
	req.Header.Set("X-Master-Key", c.MasterKey) // Обязательно замените на реальный ключ
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return fmt.Errorf("Ошибка выполнения метода DELETE: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Сервер вернул ошибку: %d, тело: %s", resp.StatusCode, body)
	}
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("Ошибка чтения ответа метода DELETE: %w", err)
	// }
	// var result JsonBinResponse
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	return fmt.Errorf("Ошибка парсинга json: %w", err)
	// }
	return nil
}

func (c *Client) PostBin(myData []byte, binName string, private bool) (*JsonBinResponse, error) {
	url := fmt.Sprintf("%s/b", baseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(myData))
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода POST: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.MasterKey)
	req.Header.Set("X-Bin-Name", binName)
	if private {
		req.Header.Set("X-Bin-Private", "true")
	} else {
		req.Header.Set("X-Bin-Private", "false")
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения метода POST: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Сервер вернул ошибку: %d, тело: %s", resp.StatusCode, body)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения ответа метода POST: %w", err)
	}
	var response JsonBinResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	fmt.Printf("Статус: %v\n", resp.Status)
	fmt.Printf("Ответ сервера: %v\n", string(body))
	return &response, nil
}

func (c *Client) GetBin(id string) (*JsonBinResponse, error) {
	url := fmt.Sprintf("%s/b/%s/latest", baseUrl, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода GET: %w", err)
	}
	req.Header.Set("X-Master-key", c.MasterKey)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения GET: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("Сервер вернул ошибку: %d, тело: %s", res.StatusCode, body)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения ответа GET: %w", err)
	}
	var result JsonBinResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга json (GET): %w", err)
	}
	return &result, nil
}

func (c *Client) UpdateBin(data []byte, id string) (*JsonBinResponse, error) {
	url := fmt.Sprintf("%s/b/%s", baseUrl, id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода PUT: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.MasterKey)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания метода PUT: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Сервер вернул ошибку: %d, тело: %s", resp.StatusCode, body)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения ответа (PUT): %w", err)
	}
	var result JsonBinResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга json (PUT): %w", err)
	}
	return &result, nil
}

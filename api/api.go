// Package api for work with json.bin
package api

import (
	"dz/bingo/config"
	"fmt"
)

func NewApi() {
	key := config.NewConfig()
	fmt.Println(key)
}

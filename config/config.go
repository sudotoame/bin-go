package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}
}
func (c *Config) Validate() error {
	if c.Key == "" {
		return fmt.Errorf("KEY пуста")
	}
	return nil
}

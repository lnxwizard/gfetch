package config

import (
	"errors"
	"github.com/pelletier/go-toml/v2"
	"os"
)

var Configs = &Config{}

type Config struct {
	TestString string `toml:"test_string"`
	TestInt    int    `toml:"test_int"`
	TestBool   bool   `toml:"test_bool"`
}

func (c *Config) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	err = toml.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Store(filename string) error {
	data, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

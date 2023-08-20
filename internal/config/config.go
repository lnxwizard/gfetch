package config

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var Configs = &Config{}

type Config struct {
	Title struct {
		BoldTitle   bool
		ItalicTitle bool
	} `toml:"Title"`
	/*Value struct {
		BoldValue   bool
		ItalicValue bool
	}*/
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

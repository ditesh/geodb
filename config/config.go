package config

import (
	"encoding/json"
	"os"
)

func (c *Config) Parse(path string) error {

	fd, err := os.Open(path)

	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(fd)

	if err = jsonParser.Decode(&c); err != nil {
		return err
	}

	return nil

}

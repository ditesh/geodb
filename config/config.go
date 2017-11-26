package config

import (
	"encoding/json"
	"os"
)

// Parse parses a JSON based config file into a config struct
func (c *Config) Parse(path string) error {

	fd, err := os.Open(path)

	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(fd)

	err = jsonParser.Decode(&c)
	return err

}

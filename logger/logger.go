package logger

import (
	"errors"
	"geodb/config"
	"io/ioutil"
	"log"
	"os"
)

var logger *log.Logger

// Configure configures the logger based on the LoggerConfig
func Configure(c config.LoggerConfig) error {

	if c.Type == "file" {

		fd, err := os.OpenFile(c.Path+"/geodb.log", os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			return err
		}

		logger = log.New(fd, "", log.Ldate|log.Ltime|log.Lshortfile)

	} else if c.Type == "discard" {

		logger = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)

	} else {
		return errors.New("none file logging types don't exist yet")
	}

	return nil

}

// Error outputs an error string with the given message
func Error(v ...interface{}) {
	logger.Print("ERROR:", v)
}

// Info outputs an error string with the given message
func Info(v ...interface{}) {
	logger.Print("INFO:", v)
}

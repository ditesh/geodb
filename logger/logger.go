package logger

import (
	"errors"
	"geodb/config"
	"log"
	"os"
)

var logger *log.Logger

func Configure(c config.LoggerConfig) error {

	if c.Type == "file" {

		fd, err := os.OpenFile(c.Path+"/geodb.log", os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			return err
		}

		logger = log.New(fd, "", log.Ldate|log.Ltime|log.Lshortfile)

	} else {
		return errors.New("none file logging types don't exist yet")
	}

	return nil

}

func Error(v ...interface{}) {
	logger.Print("ERROR:", v)
}

func Info(v ...interface{}) {
	logger.Print("INFO:", v)
}

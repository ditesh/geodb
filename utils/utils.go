package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

// Exists checks whether a given filesystem path exists
func Exists(path string) (bool, error) {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err

}

// Writable checks whether a given filesystem path is writable
func Writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

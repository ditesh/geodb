package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

func exists(path string, checkdir bool) (bool, error) {

	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	mode := fi.Mode()

	if checkdir {
		return mode.IsDir(), err
	}

	return mode.IsRegular(), err

}

// DirExists checks whether a given filesystem direcotyr exists
func DirExists(path string) (bool, error) {
	return exists(path, true)
}

// FileExists checks whether a given file exists
func FileExists(path string) (bool, error) {
	return exists(path, false)
}

// Writable checks whether a given filesystem path is writable
func Writable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

package utils

import (
	"geodb/structs"
	"io/ioutil"
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

// CreateTestDirs creates a set of test dirs and
func CreateTestDirs(n int, prefix string) ([]string, error) {

	retval := make([]string, n)

	for i := 0; i < n; i++ {

		tmpdir, err := ioutil.TempDir("", prefix)

		if err != nil {

			for _, dir := range retval {
				os.RemoveAll(dir)
			}
		}

		retval[i] = tmpdir

	}

	return retval, nil

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

func RunTableTests(tests []structs.TableTest, t func(format string, args ...interface{}), f func(string) bool) {

	for _, tt := range tests {

		if ok := f(tt.In); ok != tt.Out {
			t("in: %s, out: %t, got: %t", tt.In, tt.Out, ok)
		}

	}

}

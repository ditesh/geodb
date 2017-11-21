package utils

import (
	"fmt"
	"geodb/structs"
	"io/ioutil"
	"os"
)

// CreateTestDirs creates a set of test dirs
func CreateTestDirs(n int) ([]string, error) {

	retval := make([]string, n)

	for i := 0; i < n; i++ {

		tmpdir, err := ioutil.TempDir("", "geodbtest")

		if err != nil {

			for _, dir := range retval {
				if err := os.RemoveAll(dir); err != nil {
					fmt.Println("unable to remove test dirs")
				}
			}
		}

		retval[i] = tmpdir

	}

	return retval, nil

}

func RunTableTests(tests []structs.TableTest, t func(format string, args ...interface{}), f func(string) bool) {

	for _, tt := range tests {

		if ok := f(tt.In); ok != tt.Out {
			t("in: %s, out: %t, got: %t", tt.In, tt.Out, ok)
		}

	}

}

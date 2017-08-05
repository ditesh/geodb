package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestExists(t *testing.T) {

	var tests = []struct {
		in  string
		out bool
	}{
		{"", false},
		{".", true},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", true},
	}

	for _, tt := range tests {

		if ok, _ := Exists(tt.in); ok != tt.out {
			t.Error("expected \"" + tt.in + " existance to be " + strconv.FormatBool(tt.out) + " but got " + strconv.FormatBool(ok) + " instead")
		}

	}

	// Test dir setup
	if dir, err := ioutil.TempDir("", "utils-test"); err == nil {

		// Test dir cleanup
		defer func() {
			if err := os.RemoveAll(dir); err != nil {
				fmt.Println("unable to remove " + dir)
			}
		}()

		if ok, err := Exists(dir); !ok {
			t.Error("tmpdir ("+dir+") path should exist but did not:", err)
		}
	}

}

func TestWritable(t *testing.T) {

	if ok := Writable(""); ok {
		t.Error("empty string path should not be writable but did")
	}

	if ok := Writable("I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC"); ok {
		t.Error("random string path should not be writable but is")
	}

	// Test dir setup
	if dir, err := ioutil.TempDir("", "utils-test"); err == nil {

		// Test dir cleanup
		defer func() {
			if err := os.RemoveAll(dir); err != nil {
				fmt.Println("unable to remove " + dir)
			}
		}()

		if ok, err := Exists(dir); !ok {
			t.Error("tmpdir ("+dir+") path should be writable but is not:", err)
		}
	}
}

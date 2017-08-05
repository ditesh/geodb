package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExists(t *testing.T) {

	if ok, err := Exists(""); ok {
		t.Error("empty string path should not exist but did:", err)
	}

	if ok, err := Exists("/"); !ok {
		t.Error("/ should exist but did not:", err)
	}

	if ok, err := Exists("I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC"); ok {
		t.Error("random string path should not exist but did:", err)
	}

	if ok, err := Exists("../utils"); !ok {
		t.Error("utils path should exist but did not:", err)
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

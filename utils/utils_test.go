package utils

import (
	"fmt"
	"geodb/testhelpers"
	"io/ioutil"
	"os"
	"testing"
)

var e = testhelpers.Error{}

func TestFileExists(t *testing.T) {

	tmpfd, err := ioutil.TempFile("", "geodbtestfile")

	if err != nil {
		t.Error("unable to create tmpfile: " + err.Error())
	}

	tmpfile := tmpfd.Name()

	defer func() {

		err := tmpfd.Close()

		if err != nil {
			t.Error("unable to close tmpfd")
		}

		err = os.Remove(tmpfile)

		if err != nil {
			t.Errorf("unable to remove tmpfile: %s", tmpfile)
		}

	}()

	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{".", false},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", false},
		{tmpfile, true},
	}

	for k, tt := range tests {

		if ok, _ := FileExists(tt.in); ok != tt.out {
			e.Errorf(t, k, "mismatched error expectations")
		}
	}

}

func TestDirExists(t *testing.T) {

	tmpdir, err := ioutil.TempDir("", "geodbtestdir")

	if err != nil {
		t.Error("unable to create tmpdir: " + err.Error())
	}

	defer func() {
		err := os.RemoveAll(tmpdir)

		if err != nil {
			t.Errorf("unable to removall tmpdir: %s", tmpdir)
		}
	}()

	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{".", true},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", true},
		{tmpdir, true},
	}

	for k, tt := range tests {

		if ok, _ := DirExists(tt.in); ok != tt.out {
			e.Errorf(t, k, "mismatched error expectations")
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

		if ok, err := DirExists(dir); !ok {
			t.Error("tmpdir ("+dir+") path should be writable but is not:", err)
		}
	}
}

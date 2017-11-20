package utils

import (
	"fmt"
	"geodb/structs"
	"io/ioutil"
	"os"
	"testing"
)

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

	tests := []structs.TableTest{
		{"", false},
		{".", false},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", false},
		{tmpfile, true},
	}

	RunTableTests(tests, t.Errorf, func(s string) bool {
		ok, _ := FileExists(s)
		return ok
	})

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

	tests := []structs.TableTest{
		{"", false},
		{".", true},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", true},
		{tmpdir, true},
	}

	RunTableTests(tests, t.Errorf, func(s string) bool {
		ok, _ := DirExists(s)
		return ok
	})

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

func TestRunTableTests(t *testing.T) {

	tested := false

	errorF := func(format string, args ...interface{}) {
		tested = true
	}

	tests := []structs.TableTest{
		{"1", true},
		{"0", false},
	}

	RunTableTests(tests, errorF, func(in string) bool {
		return true
	})

	if tested == false {
		t.Error("error function was not called")
	}

}

func TestCreateTestDirs(t *testing.T) {

	dirs, err := CreateTestDirs(10, "utilstest")

	if err != nil {
		t.Fatal("expected no errors but received one")
	}

	for _, dir := range dirs {

		if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
			t.Errorf("Received %s but its not a valid directory", dir)
		}
	}

}

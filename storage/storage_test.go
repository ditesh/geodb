package storage

import (
	"geodb/structs"
	"geodb/utils"
	"io/ioutil"
	"os"
	"testing"
)

func TestInit(t *testing.T) {

	tmpdirs, err := utils.CreateTestDirs(2, "geodb")

	if err != nil {
		t.Fatal("unable to create test dirs")
	}

	defer func() {
		for _, dir := range tmpdirs {
			os.RemoveAll(dir)
		}
	}()

	err = os.Chmod(tmpdirs[0], 0400)

	if err != nil {
		t.Fatal("unable to set perms on test dir")
	}

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{tmpdirs[0], false},
		{tmpdirs[1], true},
	}

	for _, tt := range tests {

		if err := Init(tt.in); (err == nil) != tt.exp {
			t.Errorf("in: '%s', exp: %t, out: %t", tt.in, tt.exp, !tt.exp)
		}

	}

}

func TestInitMetadata(t *testing.T) {

	tmpdir, err := ioutil.TempDir("", "geodbtestdir")

	if err != nil {
		t.Fatal("unable to create tmpdir: " + err.Error())
	}

	tmpfd, err := ioutil.TempFile("", "geodbtestfile")

	if err != nil {
		t.Error("unable to create tmpfile: " + err.Error())
	}

	tmpfile := tmpfd.Name()

	defer func() {
		tmpfd.Close()
		os.RemoveAll(tmpdir)
		os.RemoveAll(tmpfile)
	}()

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{tmpdir, true},
		{tmpfile, false},
	}

	for _, tt := range tests {

		if _, err := InitMetadata(tt.in); (err == nil) != tt.exp {
			t.Errorf("in: '%s', exp: %t, out: %t", tt.in, tt.exp, !tt.exp)
		}

	}

}

func TestWritePoint(t *testing.T) {

	p := structs.Point{}

	retval := WritePoint(p)

	if retval != nil {
		t.Error("invalid return value")
	}

}

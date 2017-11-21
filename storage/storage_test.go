package storage

import (
	"errors"
	"geodb/structs"
	"geodb/utils"
	"io/ioutil"
	"os"
	"testing"
)

func TestInit(t *testing.T) {

	tmpdirs, err := utils.CreateTestDirs(2)

	if err != nil {
		t.Fatal("unable to create test dirs")
	}

	defer func() {
		for _, dir := range tmpdirs {
			err := os.RemoveAll(dir)

			if err != nil {
				t.Errorf("unable to remove test dir: %s", dir)
			}
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

func TestInvalidOpenFile(t *testing.T) {

	oldOpenFile := openFile
	defer func() { openFile = oldOpenFile }()

	openFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return nil, errors.New("invalid openFile call")
	}

	tmpdirs, err := utils.CreateTestDirs(1)

	if err != nil {
		t.Fatal("unable to create test dirs")
	}

	err = Init(tmpdirs[0])
	defer func() {
		for _, dir := range tmpdirs {
			err := os.RemoveAll(dir)

			if err != nil {
				t.Errorf("unable to remove test dir: %s", dir)
			}
		}
	}()

	if err == nil {
		t.Error("expected an error but got none")
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

		err := tmpfd.Close()

		if err != nil {
			t.Error("unable to close tmpfd")
		}

		err = os.RemoveAll(tmpdir)

		if err != nil {
			t.Errorf("unable to removeall: %s", tmpdir)
		}

		err = os.RemoveAll(tmpfile)

		if err != nil {
			t.Errorf("unable to removeall: %s", tmpdir)
		}

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

	dirs, err := utils.CreateTestDirs(1)

	if err != nil {
		t.Fatal("unable to create tempdir")
	}

	err = Init(dirs[0])
	defer func() {

		err := os.RemoveAll(dirs[0])

		if err != nil {
			t.Error("unable to remove tempdir")
		}

	}()

	if err != nil {
		t.Fatal("unable to initialise tempdir")
	}

	p := structs.Point{
		Lat: 10,
		Lng: 10,
		Elv: 10,
	}

	err = WritePoint(p, "{}")

	if err != nil {
		t.Error("invalid return value")
	}

}

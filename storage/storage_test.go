package storage

import (
	"errors"
	"geodb/testhelpers"
	"os"
	"testing"
)

var e = testhelpers.Error{}

func TestInit(t *testing.T) {

	h := testhelpers.Fs{T: t}
	tmpdirs, cb := h.CreateTestDirs(2)
	defer cb()

	h.Chmod(tmpdirs[0], 0400)

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{tmpdirs[0], false},
		{tmpdirs[1], true},
	}

	for k, tt := range tests {

		if err := Init(tt.in); (err == nil) != tt.exp {
			e.Errorf(t, k, "in: '%s', exp: %t, out: %t", tt.in, tt.exp, !tt.exp)
		}

	}

}

func TestInvalidOpenFile(t *testing.T) {

	h := testhelpers.Fs{T: t}

	oldOpenFile := openFile
	defer func() { openFile = oldOpenFile }()

	openFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return nil, errors.New("invalid openFile call")
	}

	tmpdirs, cb := h.CreateTestDirs(1)
	defer cb()

	if err := Init(tmpdirs[0]); err == nil {
		e.Errorf(t, 0, "expected an error but got none")
	}

}

func TestInitMetadata(t *testing.T) {

	h := testhelpers.Fs{T: t}
	tmpdir, cb := h.CreateTestDirs(1)
	defer cb()

	tmpfd, cb := h.CreateTestFiles(1)
	defer cb()

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{tmpdir[0], true},
		{tmpfd[0].Name(), false},
	}

	for k, tt := range tests {

		if _, err := InitMetadata(tt.in); (err == nil) != tt.exp {
			e.Errorf(t, k, "in: '%s', exp: %t, out: %t", tt.in, tt.exp, !tt.exp)
		}
	}
}

func TestWrite(t *testing.T) {

	fs := &testhelpers.Fs{T: t}
	fd, cb := fs.CreateTestFiles(1)
	defer cb()

	store = &Store{
		fd: fd[0],
		metadata: &Metadata{
			lastpos: 0,
		},
	}

	tests := [][]byte{
		[]byte{57, 52, 48, 54, 98, 100, 97, 56, 45, 52, 49, 98, 57, 45, 52, 0, 128, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{102, 97, 52, 101, 51, 100, 100, 54, 45, 54, 55, 98, 56, 45, 52, 0, 128, 0, 0, 24, 0, 0, 0, 128, 0, 123, 125, 0, 0, 0},
	}

	for k, v := range tests {

		if err := Write(v); err != nil {
			e.Errorf(t, k, "write failed")
		}
	}

}

func TestWritePoint(t *testing.T) {

	/*
		tmpdirs, err := utils.CreateTestDirs(1)

		if err != nil {
			t.Fatal("unable to create tempdir")
		}

		defer func() {

			err := utils.RemoveTestDirs(tmpdirs)

			if err != nil {
				t.Errorf("unable to remove testdir: %s", err.Error())
			}
		}()

		err = Init(tmpdirs[0])

		if err != nil {
			t.Fatal("unable to initialise tempdir")
		}

		p := geometry.Point{
			Lat:  10,
			Lng:  10,
			Elv:  10,
			Blob: "{}",
		}

		err = p.Write()

		if err != nil {
			t.Error("invalid return value")
		}

	*/
}

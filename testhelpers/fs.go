package testhelpers

import (
	"io/ioutil"
	"os"
)

var removeAll = os.RemoveAll
var tempDir = ioutil.TempDir

func removeTestDirs(dirs []string) error {

	for _, dir := range dirs {
		err := removeAll(dir)

		if err != nil {
			return err
		}
	}

	return nil
}

// CreateTestDirs creates a list of test dirs
func (fs *Fs) CreateTestDirs(n int) ([]string, func()) {

	tmpdirs := make([]string, n)

	for i := 0; i < n; i++ {

		tmpdir, err := tempDir("", "geodb")

		if err != nil {

			for _, dir := range tmpdirs {
				if err := removeAll(dir); err != nil {
					fs.T.Fatal("unable to remove test dirs")
				}
			}
		}

		tmpdirs[i] = tmpdir

	}

	cb := func() {

		err := removeTestDirs(tmpdirs)

		if err != nil {
			fs.T.Errorf("unable to remove testdir: %s", err.Error())
		}
	}

	return tmpdirs, cb

}

// Chmod runs chmod on a given path
func (fs *Fs) Chmod(path string, mode os.FileMode) {

	err := os.Chmod(path, 0400)

	if err != nil {
		fs.T.Fatal("unable to set perms on testdir")
	}

}

// CreateTestFiles creates a list of test files
func (fs *Fs) CreateTestFiles(n int) ([]*os.File, func()) {

	tmpfd := make([]*os.File, n)

	for i := 0; i < n; i++ {

		fd, err := ioutil.TempFile("", "geodb")

		if err != nil {
			fs.T.Error("unable to create tmpfile: " + err.Error())
		}

		tmpfd[i] = fd

	}

	cb := func() {

		for i := 0; i < n; i++ {

			err := tmpfd[i].Close()

			if err != nil {
				fs.T.Error("unable to close tmpfd")
			}
		}

	}

	return tmpfd, cb

}

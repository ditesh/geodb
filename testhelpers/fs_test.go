package testhelpers

import (
	"os"
	"testing"
)

func TestCreateTestDirs(t *testing.T) {

	fs := &Fs{
		T: t,
	}

	dirs, cb := fs.CreateTestDirs(10)
	defer cb()

	for _, dir := range dirs {

		if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
			t.Errorf("Received %s but its not a valid directory", dir)
		}
	}

}

func TestRemoveTestDirs(t *testing.T) {
	t.Skip()
}

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
		tmpfd.Close()
		os.Remove(tmpfile)
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
		os.RemoveAll(tmpdir)
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

func TestExtractBits(t *testing.T) {

	tests8 := []struct {
		in     uint8
		offset uint8
		length uint8
		out    uint32
	}{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{8, 1, 0, 0},
		{8, 2, 0, 0},
		{8, 0, 1, 0},
		{8, 0, 2, 0},
		{8, 0, 3, 0},
		{8, 0, 4, 0},
		{10, 2, 0, 0},
		{11, 2, 0, 0},
		{10, 0, 2, 0},
		{11, 0, 2, 0},
		{10, 2, 1, 0},
		{11, 2, 1, 0},
		{10, 1, 2, 0},
		{11, 1, 2, 0},
		{21, 2, 0, 0},
		{21, 2, 1, 0},
		{29, 2, 1, 0},
		{29, 1, 2, 0},
	}

	for _, tt := range tests8 {

		if bits, _ := ExtractBits(tt.in, tt.offset, tt.length); uint32(bits.(uint8)) != tt.out {

			t.Errorf("in: %d, offset: %d, length: %d, exp: %d, got: %d", tt.in, tt.offset, tt.length, tt.out, bits)

		}
	}

	tests32 := []struct {
		in     uint32
		offset uint8
		length uint8
		out    uint32
	}{
		{100000000, 7, 5, 31},
		{1171645696, 12, 4, 5},
	}

	for _, tt := range tests32 {

		if bits, _ := ExtractBits(tt.in, tt.offset, tt.length); uint32(bits.(uint32)) != tt.out {
			t.Errorf("in: %d, offset: %d, length: %d, exp: %d, got: %d", tt.in, tt.offset, tt.length, tt.out, bits)
		}
	}

	unsupportedTests := []struct {
		in     uint16
		offset uint8
		length uint8
		out    bool
	}{
		{1234, 1, 5, true},
	}

	for _, tt := range unsupportedTests {
		if _, err := ExtractBits(tt.in, tt.offset, tt.length); (err == nil) == true {
			t.Errorf("in: %d, offset: %d, length: %d, expected errors but received none", tt.in, tt.offset, tt.length)
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

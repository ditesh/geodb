package logger

import (
	"fmt"
	"geodb/config"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type writerWrapper struct {
	data []byte
}

func (w *writerWrapper) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)
	return len(p), nil
}

type tableTests struct {
	in  []string
	out bool
}

func runTests(tests []tableTests, t *testing.T) {

	var c config.LoggerConfig

	for _, tt := range tests {

		c.Type = tt.in[0]
		c.Path = tt.in[1]
		c.Level = tt.in[2]

		if err := Configure(c); (err == nil) != tt.out {

			exp := "no errors but received one"

			if err != nil {
				exp = "errors but did not receive one"
			}

			t.Errorf("Type: '%s', Path: '%s', Level: '%s', expected %s", tt.in[0], tt.in[1], tt.in[2], exp)
		}
	}

}

func TestConfigure(t *testing.T) {

	// Test dir setup
	dir, err := ioutil.TempDir("", "configtest")

	if err != nil {
		t.Fatal("unable to create temp dir")
		return
	}

	tests := []tableTests{
		{[]string{"", "", ""}, false},
		{[]string{"syslog", "", ""}, false},
		{[]string{"file", "", ""}, false},
		{[]string{"file", "", "debug"}, false},
		{[]string{"file", dir, "debug"}, true},
		{[]string{"discard", "", "debug"}, true},
	}

	runTests(tests, t)

	// Remove temporary dir
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println("unable to remove " + dir)
	}

	// Test dir setup
	dir, err = ioutil.TempDir("", "config-test-2")
	if err != nil {
		t.Fatal("unable to create temp dir")
		return
	}

	// Make temp dir unwriteable
	err = os.Chmod(dir, 0400)

	if err != nil {
		t.Fatal("unable to make temp dir unwriteable")
	}

	tests = []tableTests{
		{[]string{"file", dir, "debug"}, false},
	}
	runTests(tests, t)

	// Make temp dir writable again
	err = os.Chmod(dir, 0777)

	if err != nil {
		t.Fatal("unable to make temp dir unwriteable")
	}

	// Remove temporary dir
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println("unable to remove " + dir)
	}

}

func TestError(t *testing.T) {

	w := &writerWrapper{}
	logger = log.New(w, "", 0)

	Error("test")
	output := string(w.data)

	if output != "ERROR:[test]\n" {
		t.Errorf("expected 'ERROR:[test]\n' but received '%s' instead", output)
	}

}

func TestInfo(t *testing.T) {

	w := &writerWrapper{}
	logger = log.New(w, "", 0)

	Info("test")
	output := string(w.data)

	if output != "INFO:[test]\n" {
		t.Errorf("expected 'INFO:[test]\n' but received '%s' instead", output)
	}

}

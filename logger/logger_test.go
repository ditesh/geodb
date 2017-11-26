package logger

import (
	"geodb/config"
	"geodb/testhelpers"
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

	fs := &testhelpers.Fs{T: t}

	// Test dir setup
	dir, cb := fs.CreateTestDirs(2)
	defer cb()

	tests := []tableTests{
		{[]string{"", "", ""}, false},
		{[]string{"syslog", "", ""}, false},
		{[]string{"file", "", ""}, false},
		{[]string{"file", "", "debug"}, false},
		{[]string{"file", dir[0], "debug"}, true},
		{[]string{"discard", "", "debug"}, true},
	}

	runTests(tests, t)

	// Make temp dir unwriteable
	if err := os.Chmod(dir[1], 0400); err != nil {
		t.Fatal("unable to make temp dir unwriteable")
	}

	tests = []tableTests{
		{[]string{"file", dir[1], "debug"}, false},
	}
	runTests(tests, t)

	// Make temp dir writable again
	if err := os.Chmod(dir[1], 0777); err != nil {
		t.Fatal("unable to make temp dir unwriteable")
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

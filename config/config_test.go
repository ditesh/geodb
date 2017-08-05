package config

import (
	"testing"
)

func TestParse(t *testing.T) {

	var tests = []struct {
		in  string
		out bool
	}{
		{"", false},
		{"I43JSRnnGwzWFJn0TbIWRJW6TddKdMaspC2bENRC", false},
		{"../utils", false},
		{"../testdata/config/valid.json", true},
		{"../testdata/config/invalid-1.json", false},
		{"../testdata/config/invalid-2.json", false},
	}

	c := &Config{}

	for _, tt := range tests {

		if err := c.Parse(tt.in); (err == nil) != tt.out {

			if tt.out == true {
				t.Error("expected parsing \"" + tt.in + " to be valid but received an error")
			} else {
				t.Error("expected parsing \"" + tt.in + " to be invalid but did not receive an error")
			}
		}

	}

}

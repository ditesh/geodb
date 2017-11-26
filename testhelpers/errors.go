package testhelpers

import "testing"

// Errorf is a wrapper around t.Errorf
// It provides the ability to indicate which table test failed
func (e Error) Errorf(t *testing.T, k int, msg string, args ...interface{}) {
	t.Errorf("Test %d: "+msg, k, args)
}

// Fatalf is a wrapper around t.Fatalf
// It provides the ability to indicate which table test failed
func (e Error) Fatalf(t *testing.T, k int, msg string, args ...interface{}) {
	t.Fatalf("Test %d: "+msg, k, args)
}

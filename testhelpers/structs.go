package testhelpers

import "testing"

// Fs is a testhelper struct which encapsulates t and provifes fs test helpers
type Fs struct {
	T *testing.T
}

// Error is a testhelper struct which provides error test helpers
type Error struct{}

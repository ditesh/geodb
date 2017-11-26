package storage

import (
	"os"
	"sync"
)

// Store is a wrapper around key storage types
type Store struct {
	metadata *Metadata
	fd       *os.File
	datadir  string
	sync.RWMutex
}

// Metadata is a wrapper around storage metadata
type Metadata struct {
	lastpos uint64
	sync.RWMutex
}

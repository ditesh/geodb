package storage

import (
	"os"
	"sync"
)

type Store struct {
	metadata *Metadata
	fd       *os.File
	datadir  string
	sync.RWMutex
}

type Metadata struct {
	lastpos uint64
	sync.RWMutex
}

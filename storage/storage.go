package storage

import (
	"errors"
	"geodb/structs"
	"geodb/utils"
	"os"
)

var store *Store
var openFile = os.OpenFile

func writeData(data []byte) error {

	store.Lock()
	defer func() {
		store.Unlock()
	}()

	written, err := store.fd.Write(data)
	store.metadata.lastpos += uint64(written)

	return err

}

func WritePoint(p structs.Point, blob string) error {

	pr, err := NewPointRecord(p, blob)

	if err != nil {
		return err
	}

	if err := writeData(pr.Prep()); err != nil {
		return err
	}

	return nil

}

func Init(path string) error {

	metadata, err := InitMetadata(path)

	if err != nil {
		return err
	}

	// O_CREATE: create a file if non-exists
	// O_RDWR: open the file in read-write mode
	fd, err := openFile(path+"/data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)

	if err != nil {
		return err
	}

	pos, err := fd.Seek(0, 0)

	if err != nil {
		return err
	}

	metadata.lastpos = uint64(pos)

	store = &Store{
		fd:       fd,
		datadir:  path,
		metadata: metadata,
	}

	return nil

}

func InitMetadata(path string) (*Metadata, error) {

	ok, _ := utils.DirExists(path)

	if ok && !utils.Writable(path) {
		return nil, errors.New(path + " is not writable")
	} else if !ok {

		if err := os.Mkdir(path, 644); err != nil {
			return nil, err
		}
	}

	metadata := &Metadata{
		lastpos: 0,
	}

	return metadata, nil

}

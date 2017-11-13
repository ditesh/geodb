package storage

import (
	"errors"
	"fmt"
	"geodb/structs"
	"geodb/utils"
	"os"
)

var store *Store

func walWrite(data string) error {
	return errors.New("not implemented")
}

func writeData(p structs.Point) (int, error) {

	store.Lock()
	defer func() {
		store.Unlock()
	}()

	data := encodePoint(p)
	return store.fd.Write(data)

}

func updateIdx(p structs.Point) error {
	return errors.New("not implemented")
}

func updateCaches(p structs.Point) error {
	return errors.New("not implemented")
}

func WritePoint(p *structs.Point) error {

	return nil /*
		// We love locks :allthethings:
		store.Lock()

		// Serialize point to binary encoding
		data, err := serializePoint(p) // serialize to binary

		if err != nil {
			return err
		}

		// write to WAL
		if err := walWrite(data); err != nil {
			return err
		}

		if err := dataWrite(data); err != nil {
			return err
		}

		// TODO: we need to figure out how to write to index
		if err := idxWrite(p); err != nil {
			return err
		}

		if err := updateCaches(p); err != nil {
			return err
		}

		store.Unlock()
		return nil
	*/
}

func validate(path string) error {

	if ok, err := utils.DirExists(path); !ok {
		return err
	}

	if ok := utils.Writable(path); !ok {
		return errors.New("unable to write to path")
	}

	return nil

}

func Init(path string) error {

	metadata, err := InitMetadata(path)

	if err != nil {
		return err
	}

	fd, err := os.OpenFile("data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)

	if err != nil {
		return err
	}

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

		fmt.Println("Zef")
		if err := os.Mkdir(path, 644); err != nil {
			return nil, err
		}
	}

	metadata := &Metadata{
		lastpos: 0,
	}

	return metadata, nil

}

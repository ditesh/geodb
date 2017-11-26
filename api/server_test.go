package api

import (
	"geodb/config"
	"geodb/logger"
	"geodb/storage"
	"geodb/testhelpers"
	"testing"

	context "golang.org/x/net/context"
)

func TestListen(t *testing.T) {

	err := logger.Configure(config.LoggerConfig{
		Type: "discard",
	})

	if err != nil {
		t.Fatal("unable to configure logger")
	}

	if lis, err := listen(-1); err == nil {

		err := lis.Close()

		if err != nil {
			t.Fatal("unable to close listener")
		}

		t.Error("expected an error but didn't receive one")

	}

	if lis, err := listen(1); err == nil {

		err := lis.Close()

		if err != nil {
			t.Fatal("unable to close listener")
		}

		t.Error("expected an error but didn't receive one")

	}

	// Random high port
	lis, err := listen(12311)

	if err != nil {
		t.Error("did not expect an error but received one")
	} else {

		err := lis.Close()

		if err != nil {
			t.Fatal("unable to close listener")
		}

	}
}

func TestWrite(t *testing.T) {

	fs := &testhelpers.Fs{
		T: t,
	}

	dirs, cb := fs.CreateTestDirs(1)
	defer cb()

	if err := storage.Init(dirs[0]); err != nil {
		t.Fatal("unable to initialise tempdir")
	}

	p := &Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	wr := &WritePointRequest{
		P:    p,
		Blob: "test",
	}

	s := &wrapper{}

	if _, err := s.WritePoint(context.Background(), wr); err != nil {
		t.Error("unable to write point")
	}
}

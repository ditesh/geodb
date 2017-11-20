package api

import (
	"testing"

	context "golang.org/x/net/context"
)

func TestStart(t *testing.T) {

	s := &Server{
		Config{
			Port: 0,
		},
	}

	s.Start()

}

func TestWrite(t *testing.T) {

	p := &Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	wr := &WriteRequest{
		P:    p,
		Blob: "test",
	}

	s := &wrapper{}

	if _, err := s.Write(context.Background(), wr); err != nil {
		t.Error("unable to write point")
	}
}

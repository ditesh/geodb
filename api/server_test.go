package api

import (
	"testing"

	context "golang.org/x/net/context"
)

func TestWrite(t *testing.T) {

	t.Skip()

	wr := &WriteRequest{
		P: {
			Lat: 0,
			Lng: 0,
			Elv: 0,
		},
		Blob: "test",
	}

	if err := Write(context.Background(), wr); err != nil {
		t.Error("unable to write point")
	}

}

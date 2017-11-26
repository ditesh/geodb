package geometry

import (
	"bytes"
	"errors"
	"geodb/testhelpers"
	"testing"

	"github.com/google/uuid"
)

var e = testhelpers.Error{}

func TestNewPointRecord(t *testing.T) {

	t.Parallel()

	tests := []struct {
		lat  int32
		lng  int32
		elv  int32
		blob string
		len  int32
		exp  bool
	}{
		{0, 0, 0, "", 0, true},
		{1, 1, 1, "{}", 2, true},
		{91000000, -1, -1, "{'key': 'val'}", 14, false},   // invalid lat
		{-91000000, -1, -1, "{'key': 'val'}", 14, false},  // invalid lat
		{181000000, -1, -1, "{'key': 'val'}", 14, false},  // invalid lng
		{-181000000, -1, -1, "{'key': 'val'}", 14, false}, // invalid lng
		{-1, -1, -1, "{'key': 'val'}", 14, false},         // invalid elv
		{-1, -1, -1, "{'key': 'val'}", 14, false},         // invalid elv
	}

	var p Point

	for k, tt := range tests {

		p.Lat = tt.lat
		p.Lng = tt.lng
		p.Elv = tt.elv
		p.Blob = tt.blob

		pr, err := newRecord(p)

		if (err == nil) != tt.exp {
			e.Errorf(t, k, "lat: %d, lng: %d, elv: %d, exp err: %t", tt.lat, tt.lng, tt.elv, tt.exp)
		}

		if err == nil && int32(len(pr.blob)) != tt.len {
			e.Errorf(t, k, "expected bloblen: %d, received: %d", tt.len, len(pr.blob))
		}

	}

}

func TestFailedUUIDMarshaling(t *testing.T) {

	OldMarshalUUID := marshalUUID
	defer func() {
		marshalUUID = OldMarshalUUID
	}()

	marshalUUID = func(uuid uuid.UUID) ([]byte, error) {
		return nil, errors.New("error")
	}

	p := Point{
		Lat:  1,
		Lng:  1,
		Elv:  1,
		Blob: "{}",
	}

	_, err := newRecord(p)

	if err == nil {
		e.Errorf(t, 0, "expected error but received none")
	}

}

func TestSerializationPath(t *testing.T) {

	tests := []struct {
		lat  int32
		lng  int32
		elv  int32
		blob string
	}{
		{0, 0, 0, ""},
		{1, 1, 1, "{}"},
	}

	for k, tt := range tests {

		p := Point{
			Lat:  tt.lat,
			Lng:  tt.lng,
			Elv:  tt.elv,
			Blob: tt.blob,
		}

		pr, err := newRecord(p)

		if err != nil {
			e.Fatalf(t, k, "unexpected error")
		}

		serialized := toBytes(pr)
		newpr := fromBytes(serialized)

		if !bytes.Equal(pr.uuid, newpr.uuid) {
			e.Errorf(t, 0, "uuid doesn't match")
		}

		if !bytes.Equal(pr.point, newpr.point) {
			e.Errorf(t, 0, "point doesn't match")
		}

		if !bytes.Equal(pr.blob, newpr.blob) {
			e.Errorf(t, 0, "blob doesn't match")
		}
	}

}

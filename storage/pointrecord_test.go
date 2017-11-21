package storage

import (
	"geodb/structs"
	"testing"
)

func TestNewPointRecord(t *testing.T) {

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

	var p structs.Point

	for _, tt := range tests {

		p.Lat = tt.lat
		p.Lng = tt.lng
		p.Elv = tt.elv

		pr, err := NewPointRecord(p, tt.blob)

		if (err == nil) != tt.exp {
			t.Errorf("lat: %d, lng: %d, elv: %d, exp err: %t", tt.lat, tt.lng, tt.elv, tt.exp)
		}

		if pr != nil && int32(len(pr.blob)) != tt.len {
			t.Errorf("Expected bloblen: %d, received: %d", tt.len, len(pr.blob))
		}

	}

}

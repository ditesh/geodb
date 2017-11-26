package geometry

import (
	"github.com/google/uuid"
)

var marshalUUID = func(uuid uuid.UUID) ([]byte, error) {
	return uuid.MarshalBinary()
}

// Ondisk structure of PointRecord
// Byte 1-16: UUID
// Byte 17-25: Point
// Byte 26-28: Blob length
// Byte 29-xxx: Blob

// toBytes converts a PointRecord to a byte array
func toBytes(pr PointRecord) []byte {

	retval := make([]byte, 25+len(pr.blob))

	copy(retval[0:16], pr.uuid)
	copy(retval[16:25], pr.point)
	copy(retval[25:], pr.blob)

	return retval

}

// fromBytes converts a byte array to a PointRecord
func fromBytes(data []byte) PointRecord {

	return PointRecord{
		uuid:  data[0:16],
		point: data[16:25],
		blob:  data[25:],
	}

}

// newRecord creates a PointRecord representation of a Point
func newRecord(p Point) (PointRecord, error) {

	uuid := uuid.New()
	uuidbin, err := marshalUUID(uuid)

	if err != nil {
		return PointRecord{}, err
	}

	// Serialize point to a sequence of bytes
	pointBytes, err := p.Encode()

	if err != nil {
		return PointRecord{}, err
	}

	return PointRecord{
		uuid:  uuidbin,
		point: pointBytes,
		blob:  []byte(p.Blob),
	}, nil

}

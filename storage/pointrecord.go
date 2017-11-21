package storage

import (
	"encoding/binary"
	"geodb/structs"

	"github.com/google/uuid"
)

// Ondisk structure of PointRecord
// Byte 1-16: UUID
// Byte 17-25: Point
// Byte 26-28: Blob length
// Byte 29-xxx: Blob
func (pr *PointRecord) Prep() []byte {

	retval := make([]byte, 28+len(pr.blob))
	copy(retval[0:15], pr.uuid)
	copy(retval[16:24], pr.point)
	copy(retval[25:27], pr.bloblen)
	copy(retval[28:], pr.blob)

	return retval

}

func NewPointRecord(p structs.Point, blob string) (*PointRecord, error) {

	uuid := uuid.New()

	// Serialize point to a sequence of bytes
	pointBytes, err := encodePoint(p)

	if err != nil {
		return nil, err
	}

	pr := &PointRecord{
		uuid:  []byte(uuid.String()),
		point: pointBytes,
		blob:  []byte(blob),
	}

	pr.bloblen = make([]byte, 3)
	binary.PutUvarint(pr.bloblen, uint64(len(pr.blob)))
	return pr, nil

}

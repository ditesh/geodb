package geometry

// Point represents a latlngelv with associated blob
type Point struct {
	Lat  int32
	Lng  int32
	Elv  int32
	Blob string
}

// PointRecord is a collection of tuples that represent a point record
type PointRecord struct {
	uuid  []byte
	blob  []byte
	point []byte
}

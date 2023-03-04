package mapper

type XMLMapper interface {
	Decode(origin []byte, destination interface{}) error
}

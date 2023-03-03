package adapter

type XMLMapper interface {
	Decode(origin []byte, destination interface{}) error
}

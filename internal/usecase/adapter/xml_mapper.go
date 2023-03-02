package adapter

type XMLMapper interface {
	MapXMLToModel(origin []byte, destination interface{}) error
}

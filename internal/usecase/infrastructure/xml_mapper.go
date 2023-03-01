package infrastructure

type XMLMapper interface {
	MapXMLToModel(xml string, model interface{}) error
}

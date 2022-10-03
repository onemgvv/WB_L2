package pkg

type AnalyticDataType string

const (
	XmlDataT AnalyticDataType  = "xml_data"
	JsonDataT 									 = "json_data"
)

type AnalyticLoader interface {
	GetType() AnalyticDataType
	Accept(Visitor)
}
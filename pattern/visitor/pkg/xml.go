package pkg

type XmlData struct {
	Data string
}

func NewXmlData(data string) *XmlData {
	return &XmlData{data}
}

func (x *XmlData) GetType() AnalyticDataType {
	return XmlDataT
}

func (x *XmlData) Accept(v Visitor) {
	v.VisitForXML(x)
}

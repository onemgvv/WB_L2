package pkg

type Visitor interface {
	VisitForXML(*XmlData)
	VisitForJSON(*JsonData)
}
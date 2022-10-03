package pkg

type JsonData struct {
	Data string
}

func NewJsonData(data string) *JsonData {
	return &JsonData{data}
}


func (j *JsonData) GetType() AnalyticDataType {
	return JsonDataT
}

func (j *JsonData) Accept(v Visitor) {
	v.VisitForJSON(j)
}

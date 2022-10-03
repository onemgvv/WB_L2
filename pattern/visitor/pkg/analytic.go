package pkg

type AnalyticData struct {
	Count int
	Name string
}

func NewAnalyticData(count int, name string) *AnalyticData {
	return &AnalyticData{count, name}
}
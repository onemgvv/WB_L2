package pkg

import (
	"fmt"
	"time"
)

type LoadData struct {
}

func NewLoadData() *LoadData {
	return &LoadData{}
}

func (l *LoadData) VisitForXML(x *XmlData) {
	fmt.Println("Loading xml data to Analytics server")
	fmt.Println(x.Data)
	time.Sleep(time.Second * 2)
	fmt.Println("XML data Loaded successfuly")
}

func (l *LoadData) VisitForJSON(j *JsonData) {
	fmt.Println("Loading json data to Analytics server")
	fmt.Println(j.Data)
	time.Sleep(time.Second * 2)
	fmt.Println("JSON data Loaded successfuly")
}
package visitor

import (
	"fmt"
	"patterns/visitor/pkg"
)

func Visitor() {
	xml := pkg.NewXmlData("<xml></xml>")
	json := pkg.NewJsonData("{\"key\": \"value\"}")

	dataLoader := pkg.NewLoadData()

	xml.Accept(dataLoader)
	fmt.Println()
	json.Accept(dataLoader)
}
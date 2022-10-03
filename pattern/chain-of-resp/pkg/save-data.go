package pkg

import "fmt"

type SaveDataService struct {
	Next Service
}

func (sds *SaveDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("Data not update.")
		return
	}
	fmt.Println("Save data")
}

func (sds *SaveDataService) SetNext(svc Service) {
	sds.Next = svc
}
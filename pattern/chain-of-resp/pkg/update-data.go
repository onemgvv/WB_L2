package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (ups *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] already update.\n", ups.Name)
		ups.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from device [%s]\n", ups.Name)
	data.UpdateSource = true
	ups.Next.Execute(data)
}

func (ups *UpdateDataService) SetNext(svc Service) {
	ups.Next = svc
}
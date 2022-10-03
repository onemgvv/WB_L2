package chain_of_resp

import "patterns/chain-of-resp/pkg"

func ChainOfResp() {
	device := pkg.Device{Name: "Device-1"}
	uds := pkg.UpdateDataService{Name: "Update-1"}
	sds := pkg.SaveDataService{}

	device.SetNext(&uds)
	uds.SetNext(&sds)

	data := pkg.Data{}

	device.Execute(&data)
}
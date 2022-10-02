package factory

import "patterns/factory/pkg"

var transports = []pkg.DeliveryType{pkg.Ground, pkg.Air, pkg.Water, "underground"}

func Factory() {
	for _, t := range transports {
		tr := pkg.NewTransport(t)
		if tr == nil {
			continue
		}
		tr.Delivery()
	}
}

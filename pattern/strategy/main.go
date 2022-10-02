package strategy

import "patterns/strategy/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{&pkg.FootStrategy{}, &pkg.PrivateStrategy{}, &pkg.PublicStrategy{}}
)

func Strategy() {
	nav := pkg.Navigator{}

	for _, s := range strategies {
		nav.SetStrategy(s)
		nav.Route(start, end)
	}
}

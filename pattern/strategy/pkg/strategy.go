package pkg

type Strategy interface {
	Route(start, end int)
}
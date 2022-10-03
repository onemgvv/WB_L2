package pkg

import "fmt"

type Scooter struct {
	taken bool
}

func (s *Scooter) take() {
	s.taken = true
	fmt.Println("[Index GO] Scooter is taken.")
}

func (s *Scooter) put() {
	s.taken = false
	fmt.Println("[Index GO] Scooter is put back. Rent price 10$")
}
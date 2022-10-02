package main

import (
	"fmt"
	"log"
)

func main() {
	str := "a4bc2d5e"
	res, err := Unpack(str)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

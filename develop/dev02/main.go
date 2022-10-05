package main

import (
	"fmt"

	"github.com/onemgvv/WB_L2/develop/dev01/pkg"
)

func main() {
	str := "a4bc2d5e"
	res := pkg.Unpack(str)

	fmt.Println(res)
}

package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func Unpack(str string) (result string){
	// var ekr bool
	if str == "" {
		return ""
	}

	if _, err := strconv.Atoi(str); err == nil {
		return "некорректная строка"
	}

	symbs := []rune(str)
	fmt.Println(len(symbs))
	last := symbs[0]

	// проходим в цикле по массиву
	for i := 0; i < len(symbs); i++ {
		if num, err := strconv.Atoi(string(symbs[i])); err == nil {
			result += strings.Repeat(string(last), num-1)
			last = symbs[i+1]
		} else {
			result += string(symbs[i])
			last = symbs[i]
		}
	}

	return
}

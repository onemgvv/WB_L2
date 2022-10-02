package main

import (
	"errors"
	"strconv"
	"strings"
)

func Unpack(str string) (result string, err error) {
	if str == "" {
		return "", nil
	}

	if _, err = strconv.Atoi(str); err == nil {
		return "", errors.New("некорректная строка")
	}

	symbs := []rune(str)
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

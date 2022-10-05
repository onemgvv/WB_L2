package pkg

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	mock := [][]string{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", "некорректная строка"},
		{"", ""},
	}

	for _, v := range mock {
    res := Unpack(v[0])
		if res != v[1] {
			t.Errorf("[ERROR]: Want %s, Got %s", v[1], res)
		}
  }
}
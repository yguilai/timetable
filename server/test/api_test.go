package test

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tt := "[黄卓2]"
	te := []rune(tt)

	fmt.Println(string(te[1 : len(te)-1]))

}

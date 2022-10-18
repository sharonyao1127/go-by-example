package main

import (
	"fmt"
)

func main() {
	//fmt.Println(bracket("{"))
	fmt.Println(fola(6))

}

func bracket(str string) bool {

	str2 := []uint8{}

	for i := range str {
		str2 = append(str2, str[i])
	}

	for len(str2) > 0 {
		//最右
		v := str2[len(str2)-1]

		var val uint8

		switch v {
		case '}':
			val = '{'
		case ']':
			val = '['
		case ')':
			val = '('
		}

		if str2[0] == val {
			str2 = str2[1 : len(str2)-1]
			if len(str2) == 0 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func fola(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	ret := fola(n-1) + fola(n-2)
	return ret
}

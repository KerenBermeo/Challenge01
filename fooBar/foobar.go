package foobar

import "strconv"

func fooBar(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Foo Bar"
	} else if num%3 == 0 {
		return "Foo"
	} else if num%5 == 0 {
		return "Bar"
	} else {
		change := strconv.Itoa(num)
		return change
	}
}

func RangeNum(x int) string {
	resultado := ""
	for i := 1; i < x; i++ {
		resultado = resultado + fooBar(i) + "->"
	}
	return resultado
}

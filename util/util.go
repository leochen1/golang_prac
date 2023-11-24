package util

import "fmt"

var i = 0
var F = func(s string) int {
	fmt.Printf("本次被%s\n", s)
	i++
	fmt.Printf("匿名工具函數被第%v次調用\n", i)
	return i
}





